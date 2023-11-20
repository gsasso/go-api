package main

import (
	"github.com/gsasso/go-api/internal/app/json_api"
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
	server := json_api.APICustomerServer(":3000")
	server.Run()
}
