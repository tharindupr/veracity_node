package main

import (
	"veracity_node/internal/routes"
)

func main() {
	// Load environment variables and connect to the database
	// initializers.LoadEnvVariables()
	// initializers.ConnectToDB()

	// Set up and start the router
	r := router.SetupRouter()
	r.Run(":8080")
}
