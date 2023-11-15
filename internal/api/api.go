package api

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gsasso/go-api/internal/logging"
	"github.com/gsasso/go-api/internal/service"
	"github.com/gsasso/go-api/internal/types"
)

type apiFunc func(context.Context, http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(apiFn apiFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	svc        service.CustomerFetcher
}

func APICustomerServer(listenAddr string) *APIServer {
	customerFetcher := logging.NewLogService(&service.CustomerFetcherService{})
	return &APIServer{
		listenAddr: listenAddr,
		svc:        customerFetcher,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", makeHTTPHandlerFunc(s.handleCustomer))
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetCustomer(ctx, w, r)
	case http.MethodPost:
		return s.handlePostCustomer(ctx, w, r)
	default:
		return fmt.Errorf("Method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("Id")
	var customReq types.CustomerRequest
	//Func to wrap the parameters into our CustomerReq?
	customReq.Id = id

	customerResponse, _ := s.svc.FetchCustomer(ctx, customReq)
	return WriteJson(w, http.StatusOK, &customerResponse)
}

func (s *APIServer) handlePostCustomer(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, &types.CustomerResponse{})
}
