package controller

import (
	"context"
	"math/rand"
	//handler "github.com/gsasso/go-api/internal/generated/v1/v1connect"
)

type CustomerHandler struct {
	//handler.UnimplementedCustomerIntelligenceHandler
}

func NewCustomerController() *CustomerHandler {
	return &CustomerHandler{}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	reqid := rand.Intn(10000)
	ctx = context.WithValue(ctx, "requestID", reqid)
	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}

	return resp, err
}
