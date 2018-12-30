package grpc

import (
	"log"
	"net"

	ddb "github.com/mbmanthey/foodtruck/dynamo"
	mdb "github.com/mbmanthey/foodtruck/mongodb"
	pb "github.com/mbmanthey/foodtruck/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port  = ":50052"
	local = false
)

func NewServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if local {
		session, err := mdb.CreateSession("localhost:27017")
		if err != nil {
			log.Fatalf("Faled to create mongodb session: %v", err)
		}
		server := grpc.NewServer()
		pb.RegisterTruckServiceServer(server, &mdb.Service{Session: session})
		// Register reflection service on gRPC server.
		reflection.Register(server)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	} else {
		session, err := ddb.CreateSession()
		if err != nil {
			log.Fatalf("Faled to create DynamoDB session: %v", err)
		}
		server := grpc.NewServer()
		pb.RegisterTruckServiceServer(server, &ddb.Service{Session: session})
		// Register reflection service on gRPC server.
		reflection.Register(server)
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}

}

func main() {
	NewServer()
}
