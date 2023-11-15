package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gsasso/go-api/internal/types"
)

var customerMocks = map[string]string{
	"001": "131",
	"002": "131423",
	"003": "145131",
}

func MockCustomerFetcher(ctx context.Context, customreq types.CustomerRequest) (types.CustomerResponse, error) {
	time.Sleep(100 * time.Microsecond)
	var customResponse types.CustomerResponse
	gedi, ok := customerMocks[customreq.Id]
	if !ok {
		return customResponse, fmt.Errorf("Customer not found with Id %s", customreq.Id)
	}
	customResponse.Id = customreq.Id
	customResponse.GEDI = gedi
	return customResponse, nil
}
