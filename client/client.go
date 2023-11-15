package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gsasso/go-api/internal/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) FetchCustomer(ctx context.Context, customReq types.CustomerRequest) (*types.CustomerResponse, error) {
	req, err := http.NewRequest("get", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	customerResponse := new(types.CustomerResponse)
	if err := json.NewDecoder(resp.Body).Decode(customerResponse); err != nil {
		return nil, err
	}

	return customerResponse, nil
}
