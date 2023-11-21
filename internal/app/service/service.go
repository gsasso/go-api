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
type CustomerFetcherService struct{}

func (s *CustomerFetcherService) FetchCustomer(ctx context.Context, req types.CustomerRequest) (types.CustomerResponse, error) {
	//Call DB in storage.go and if not found call bisnode API
	return MockCustomerFetcher(ctx, req)
}
