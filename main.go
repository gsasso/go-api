package main

import (
	"context"
	"fmt"
)

func main() {

	//Initializing the service
	svc := newLogService(&customerFetcher{})

	//Using the service straight forward!
	customer, _ := svc.FetchCustomer(context.Background(), "001")

	fmt.Println(customer)

	//Initializing the Server
	server := APICustomerServer(":3000", svc)
	server.Run()
}
