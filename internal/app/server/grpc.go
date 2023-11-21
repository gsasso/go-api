package server

import (
	"net"

	"github.com/gsasso/go-api/internal/app/controller"
	v1 "github.com/gsasso/go-api/internal/generated/v1"
	"google.golang.org/grpc"
)

type CustomerIntelligenceServer struct {
	server *grpc.Server
}

func makeGRPCServer(c controller.CustomerController) *CustomerIntelligenceServer {

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	v1.RegisterCustomerIntelligenceServer(server, c)

	return &CustomerIntelligenceServer{
		server: server,
	}
}

func (my *CustomerIntelligenceServer) Start() error {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}

	my.server.Serve(listener)

	return nil

}
