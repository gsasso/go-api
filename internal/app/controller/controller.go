package controller

import (
	"context"

	"github.com/gsasso/go-api/internal/app/service"
	"github.com/gsasso/go-api/internal/app/types"
	v1 "github.com/gsasso/go-api/internal/generated/v1"
)

type CustomerController struct {
	v1.UnimplementedCustomerIntelligenceServer
	svc service.CustomerIntelligence
}

func NewCustomerIntelligenceController(svc service.CustomerIntelligence) *CustomerController {
	return &CustomerController{
		svc: svc,
	}
}

func (s *CustomerController) FetchCustomer(ctx context.Context, req *v1.CustomerFetchRequest) (*v1.CustomerFetchResponse, error) {
	resp, err := s.svc.FetchCustomer(ctx, types.CustomerRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &v1.CustomerFetchResponse{Id: resp.Id, GEDI: resp.GEDI}, nil
}
