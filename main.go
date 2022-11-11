package main

import (
	"log"
	"net/http"

	petstore "github.com/santoshanand/ogen-go-demo/petstore"
	"github.com/santoshanand/ogen-go-demo/services"
)

func main() {
	// Create service instance.
	service := &services.PetsService{
		Pets: map[int64]petstore.Pet{},
	}
	// Create generated server.
	srv, err := petstore.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":5000", srv); err != nil {
		log.Fatal(err)
	}
}
