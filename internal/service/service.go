package service

import (
	"context"

	"github.com/gsasso/go-api/internal/types"
)

//Interface that can fetch a customer
type CustomerFetcher interface {
	FetchCustomer(context.Context, types.CustomerRequest) (types.CustomerResponse, error)
}

//Implement CustomerFetcher interface
type CustomerFetcherService struct{}

func (s *CustomerFetcherService) FetchCustomer(ctx context.Context, customreq types.CustomerRequest) (types.CustomerResponse, error) {
	//Call DB in storage.go and if not found call bisnode API
	return MockCustomerFetcher(ctx, customreq)
}
