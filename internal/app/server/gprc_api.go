package server

import (
	"github.com/gsasso/go-api/internal/app/controller"
	"google.golang.org/grpc"
)

type CustomerIntelligenceServer struct {
	server *grpc.Server
}

func makeGRPCServer(c controller.CustomerController) *CustomerIntelligenceServer {

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	//v1connect.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	return &CustomerIntelligenceServer{
		server: server,
	}
}

// type GRPCPriceFetcherServer struct {
// 	svc PriceService
// 	proto.UnimplementedPriceFetcherServer
// }

// func NewGRPCPriceFetcherServer(svc PriceService) *GRPCPriceFetcherServer {
// 	return &GRPCPriceFetcherServer{
// 		svc: svc,
// 	}
// }

// func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
// 	reqid := rand.Intn(10000)
// 	ctx = context.WithValue(ctx, "requestID", reqid)
// 	price, err := s.svc.FetchPrice(ctx, req.Ticker)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := &proto.PriceResponse{
// 		Ticker: req.Ticker,
// 		Price:  float32(price),
// 	}

// 	return resp, err
// }
