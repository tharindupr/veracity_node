package main

import (
	"veracity_node/internal/routes"
	"veracity_node/config"
	"fmt"
	"os"
	"log"
)

func main() {
	// Load environment variables and connect to the database
	// initializers.LoadEnvVariables()
	// initializers.ConnectToDB()

	// Set up and start the router

	// Set a custom config path via environment variable or default to "./config"
	configPath := "./config"
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		configPath = path
	}

	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Use the configuration
	fmt.Printf("Server running on %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("CA Certificate Path: %s\n", cfg.CACertPath)
	// Print the content of the CA certificate
	fmt.Println("CA Certificate Content:")
	fmt.Println(cfg.CACertContent)

	r := router.SetupRouter()
	r.Run(":8080")
}
