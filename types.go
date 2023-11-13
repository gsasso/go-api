package main

import "time"

type CustomerResponse struct {
	Id        string    `json:"ID"`
	GEDI      string    `json:"GEDI"`
	CreatedAt time.Time `json:"createdAt"`
}
