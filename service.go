package main

import (
	"context"
	"fmt"
	"time"
)

//Interface that can fetch a customer
type CustomerFetcher interface {
	FetchCustomer(context.Context, string) (string, error)
}

//Implement CustomerFetcher interface
type customerFetcher struct{}

func (s *customerFetcher) FetchCustomer(ctx context.Context, id string) (string, error) {
	return MockCustomerFetcher(ctx, id)
}

var customerMocks = map[string]string{
	"001": "131",
	"002": "131423",
	"003": "145131",
}

func MockCustomerFetcher(ctx context.Context, id string) (string, error) {
	time.Sleep(100 * time.Microsecond)
	customer, ok := customerMocks[id]
	if !ok {
		return customer, fmt.Errorf("Customer not found %s", customer)
	}
	return customer, nil
}
