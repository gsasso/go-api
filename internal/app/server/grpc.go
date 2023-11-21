package server

import (
	"fmt"
	"net"

	"github.com/google/wire"
	"github.com/gsasso/go-api/internal/app/controller"
	"github.com/gsasso/go-api/internal/app/service"
	proto "github.com/gsasso/go-api/internal/generated/v1"
	"google.golang.org/grpc"
)

type CustomerIntelligenceServer struct {
	server *grpc.Server
}

func ProvideGRPCServer(c *controller.CustomerController) *CustomerIntelligenceServer {

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterCustomerIntelligenceServer(server, c)

	return &CustomerIntelligenceServer{
		server: server,
	}
}

func (my *CustomerIntelligenceServer) Start() error {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	my.server.Serve(listener)
	return nil
}

var ServerProvider = wire.NewSet(service.ProvideCustomerService, wire.Bind(new(service.CustomerIntelligence), new(*service.CustomerService)),
	controller.ProvideCustomerIntelligenceController, ProvideGRPCServer)
