package main

import (
	"flag"
	"fmt"
	"log"

	"calc_service/internal/server"
)

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	address := fmt.Sprintf(":%s", *port)

	srv := server.NewServer(address)
	log.Printf("Starting server on %s", address)

	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
