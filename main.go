package main

import (
	"log"
	"net/http"

	grpc "github.com/mbmanthey/foodtruck/grpc"
	rest "github.com/mbmanthey/foodtruck/rest"
)

func main() {
	go rest.NewServer()
	go grpc.NewServer()
	http.Handle("/", http.FileServer(http.Dir("./react/build")))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
