package service

import (
	"context"

	"github.com/gsasso/go-api/internal/app/types"
)

//Interface that can fetch a customer
type CustomerIntelligence interface {
	FetchCustomer(context.Context, types.CustomerRequest) (types.CustomerResponse, error)
}

//Implement CustomerFetcher interface
type CustomerService struct{}

func (s *CustomerService) FetchCustomer(ctx context.Context, req types.CustomerRequest) (types.CustomerResponse, error) {
	//Call DB in storage.go to get GEDI and if not found call bisnode API MATCH
	return MockCustomerFetcher(ctx, req)
}

func ProvideCustomerService() *CustomerService {
	return &CustomerService{}
}
