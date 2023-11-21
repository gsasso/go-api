package controller

import (
	"context"

	"github.com/gsasso/go-api/internal/app/service"
	"github.com/gsasso/go-api/internal/app/types"
	proto "github.com/gsasso/go-api/internal/generated/v1"
)

type CustomerController struct {
	proto.UnimplementedCustomerIntelligenceServer
	svc service.CustomerIntelligence
}

func ProvideCustomerIntelligenceController(svc service.CustomerIntelligence) *CustomerController {
	return &CustomerController{
		svc: svc,
	}
}

func (s *CustomerController) FetchCustomer(ctx context.Context, req *proto.CustomerFetchRequest) (*proto.CustomerFetchResponse, error) {
	resp, err := s.svc.FetchCustomer(ctx, types.CustomerRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &proto.CustomerFetchResponse{Id: resp.Id, GEDI: resp.GEDI}, nil
}
