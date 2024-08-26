package main

import (
	"net/http"
	"veracity_node/initializers"

	"github.com/gin-gonic/gin"
)

// Structs to represent the request and response bodies
type UserTransaction struct {
	Data      map[string]interface{} `json:"data" binding:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Signature string                 `json:"signature" binding:"required"`
}

type User struct {
	ID       string                 `json:"id"`
	AssetID  string                 `json:"asset_id"`
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
	InputID  string                 `json:"input_id"`
}

// Simulated database
var users = []User{}

// func init() {
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDB()
// }

func main() {
	router := gin.Default()

	// PUT /user - Create a new user
	router.PUT("/user", createUser)

	// GET /users/:user_id - Get user by ID
	router.GET("/users/:user_id", getUserByID)

	// GET /users - Get all users
	router.GET("/users", getAllUsers)

	router.Run(":8080") // Start the server on port 8080
}

func createUser(c *gin.Context) {
	// Simulate admin user check
	admin := c.GetHeader("X-Admin")
	if admin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can create a user"})
		return
	}

	var userTransaction UserTransaction
	if err := c.ShouldBindJSON(&userTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate user validation (e.g., public key validation)
	if userTransaction.Data["public_key"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	// Simulate saving the user to a database
	newUser := User{
		ID:       "some_unique_id", // Generate a unique ID
		AssetID:  "some_asset_id",  // This could be derived from the transaction
		Data:     userTransaction.Data,
		Metadata: userTransaction.Metadata,
		InputID:  "some_input_id", // This could be linked to a previous transaction
	}

	users = append(users, newUser)

	c.JSON(http.StatusOK, newUser.ID) // Return the unique ID
}

func getUserByID(c *gin.Context) {
	userID := c.Param("user_id")

	// Simulate searching for the user in the database
	for _, user := range users {
		if user.ID == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
