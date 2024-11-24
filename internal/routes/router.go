package router

import (
	"veracity_node/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.PUT("/user", handler.CreateUser)
	// r.GET("/users/:user_id", handler.GetUserByID)
	// r.GET("/users", handler.GetAllUsers)

	return r
}
