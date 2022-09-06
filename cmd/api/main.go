package main

import (
	"github.com/tulioguaraldo/go-microservice/app/server"
	"github.com/tulioguaraldo/go-microservice/config/env"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.NewServer()
	server.RunServer()
}
