package main

import (
	"fmt"

	"github.com/gsasso/go-api/internal/app/controller"
	"github.com/gsasso/go-api/internal/app/service"
)

func main() {

	//Testing client
	// client := client.New("http://localhost:3000")
	// customerReq := types.CustomerRequest{
	// 	Id: "001",
	// }
	// customerResponse, err := client.FetchCustomer(context.Background(), customerReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(customerResponse.GEDI)

	//Bootstrap everything
	service := &service.CustomerFetcherService{}
	controller := controller.NewCustomerIntelligenceController(service)
	grpcServer := server.makeGRPCServer(controller)
	grpcServer.Start()

	if controller != nil {
		fmt.Println("controller and service initialized")
	}
	//grpcServer := server.makeGRPCServer(controller)
	//grpcServer.Start()

	//server := json_api.APICustomerServer(":3000")
	//server.Run()
}
