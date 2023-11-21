package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gsasso/go-api/internal/app/types"
)

var customerMocks = map[string]string{
	"001": "131",
	"002": "131423",
	"003": "145131",
}

func MockCustomerFetcher(ctx context.Context, req types.CustomerRequest) (types.CustomerResponse, error) {
	time.Sleep(100 * time.Microsecond)
	var customResponse types.CustomerResponse
	gedi, ok := customerMocks[req.Id]
	if !ok {
		return customResponse, fmt.Errorf("Customer not found with Id %s", req.Id)
	}
	customResponse.Id = req.Id
	customResponse.GEDI = gedi
	return customResponse, nil
}
