package main

func main() {

	//Bootstrap everything
	// service := &service.CustomerService{}
	// controller := controller.ProvideCustomerIntelligenceController(service)
	// grpcServer := server.ProvideGRPCServer(controller)
	// grpcServer.Start()

	//Bootstrap everything with Wire
	server := InitializeApp()
	server.Start()

}
