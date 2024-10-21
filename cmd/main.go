package main

import (
	"log"

	"github.com/4Sao1Sad/yuemnoi-activity/db"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/config"
)

func main() {
	// Load the configuration
	cfg := config.Load()

	// Initialize the MongoDB connection and get the database instance
	mongoDB, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize MongoDB: %v", err)
	}

	// Initialize the gRPC server and pass the MongoDB database
	err = db.ServerInit(cfg, mongoDB)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
