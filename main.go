package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/30Piraten/beerPong/config"
	router "github.com/30Piraten/beerPong/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Beer Pong Permissions Game")

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env variables: %v", err)
	}

	// Initialize Redis & Permit.io
	config.RedisInit()
	config.PermitInit()

	r := router.Router()
	fmt.Println("Server trekking at :1224")
	if err := http.ListenAndServe(":1224", r); err != nil {
		log.Fatalf("Server failed to start trekking: %v", err)
	}
}
