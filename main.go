package main

import (
	"log"
	"net"

	app "github.com/isaqueveras/authentication-microservice/application/authuser"
	config "github.com/isaqueveras/authentication-microservice/configuration"
	inter "github.com/isaqueveras/authentication-microservice/interfaces/authuser"
	gogrpc "google.golang.org/grpc"
)

func main() {
	// loading config of system
	config.Load()

	// Listen on port
	listen, err := net.Listen("tcp", config.Get().Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := gogrpc.NewServer()

	// Auth registration server
	app.RegisterAuthServer(server, &inter.Server{})

	// Message of success
	log.Println("Server is running in port", config.Get().Address)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
