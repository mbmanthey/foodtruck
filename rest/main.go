package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pb "github.com/mbmanthey/foodtruck/proto"
	"google.golang.org/grpc"
)

//NewServer will start the Rest listener.
func NewServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/truck", GetAll).Methods("GET")
	router.HandleFunc("/api/truck/{id}", GetTruck).Methods("GET")
	router.HandleFunc("/api/truck/{id}", CreateTruck).Methods("POST")
	http.Handle("/api/", router)
	if err := http.ListenAndServe(":50051", nil); err != nil {
		panic(err)
	}
}

//GRPC address
const (
	port = "localhost:50052"
)

//GetAll will return all trucks.
func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to grpc: %v", err)
	}
	defer conn.Close()
	client := pb.NewTruckServiceClient(conn)
	getAll, err := client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list trucks: %v", err)
	}
	json.NewEncoder(w).Encode(getAll.Trucks)
}

//GetTruck will return a truck from an ID.
func GetTruck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to grpc: %v", err)
	}
	defer conn.Close()
	client := pb.NewTruckServiceClient(conn)
	get, err := client.Get(context.Background(), &pb.GetRequest{ID: params["ID"]})
	if err != nil {
		log.Fatalf("Could not get truck: $v", err)
	}
	json.NewEncoder(w).Encode(get.Truck)
}

//CreateTruck will create a new truck.
func CreateTruck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to grpc: %v", err)
	}
	defer conn.Close()
	client := pb.NewTruckServiceClient(conn)
	create, err := client.Create(context.Background(), &pb.Truck{ID: params["ID"]})
	if err != nil {
		log.Fatalf("Could not create truck: %v", err)
	}
	json.NewEncoder(w).Encode(create.Truck)

}
