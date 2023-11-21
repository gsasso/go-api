package main

func main() {

	//Bootstrap everything
	// service := logging.ProvideLogService(&service.CustomerService{})
	// controller := controller.ProvideCustomerIntelligenceController(service)
	// grpcServer := server.ProvideGRPCServer(controller)
	// grpcServer.Start()

	// //Bootstrap everything with Wire
	server := Initialize()
	server.Start()

}
