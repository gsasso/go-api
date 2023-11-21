package main

import (
	"github.com/google/wire"
	"github.com/gsasso/go-api/internal/app/server"
)

func InitializeApp() *server.CustomerIntelligenceServer {

	wire.Build(server.ServerProvider)
	return &server.CustomerIntelligenceServer{}
}
