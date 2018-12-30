package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/mbmanthey/foodtruck/proto"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50052"
	defaultFilename = "test.json"
)

func parseFile(file string) (*pb.Truck, error) {
	var truck *pb.Truck
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &truck)
	return truck, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not conect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTruckServiceClient(conn)
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	truck, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse: %v", err)
	}
	truck.ID = uuid.Must(uuid.NewV4()).String()
	r, err := client.Create(context.Background(), truck)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	log.Printf("Created: %t", r.Created)
	getAll, err := client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list trucks: %v", err)
	}
	for _, v := range getAll.Trucks {
		log.Println(v)
	}
}
