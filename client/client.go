package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gsasso/go-api/internal/app/types"
	proto "github.com/gsasso/go-api/internal/generated/v1"
	"google.golang.org/grpc"
)

func GRPCClient(remoteAddr string) (proto.CustomerIntelligenceClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := proto.NewCustomerIntelligenceClient(conn)

	return c, nil
}

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) FetchCustomer(Id string) (*types.CustomerResponse, error) {
	endpoint := fmt.Sprintf("%s?Id=%s", c.endpoint, Id)
	req, err := http.NewRequest("GET", endpoint, nil)
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
