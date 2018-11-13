package main

import (
	"log"
	"net"

	db "github.com/mbmanthey/foodtruck/mongodb"
	pb "github.com/mbmanthey/foodtruck/proto"
	rest "github.com/mbmanthey/foodtruck/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go rest.NewServer()
	session, err := db.CreateSession("localhost:27017")
	if err != nil {
		log.Fatalf("Faled to create mongodb session: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterTruckServiceServer(server, &db.Service{Session: session})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
