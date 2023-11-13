package client

import "context"

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) FetchCustomer(ctx context.Context, id string)
