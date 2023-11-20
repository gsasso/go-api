package types

import "time"

type CustomerResponse struct {
	Id        string    `json:"ID"`
	GEDI      string    `json:"GEDI"`
	CreatedAt time.Time `json:"createdAt"`
}

type CustomerRequest struct {
	Id string
}
