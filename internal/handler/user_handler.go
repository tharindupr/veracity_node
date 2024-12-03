package handler

import (
	"net/http"
	"veracity_node/internal/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	admin := c.GetHeader("X-Admin")
	//ledger := c.GetHeader("Ledger")
	if admin != "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can create a user"})
		return
	}


	var userTransaction model.UserTransaction
	
	if err := c.ShouldBindJSON(&userTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userID, err := service.CreateUser(userTransaction)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, "abcdedf")
}

// func GetUserByID(c *gin.Context) {
// 	userID := c.Param("user_id")

// 	user, err := service.GetUserByID(userID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// func GetAllUsers(c *gin.Context) {
// 	users := service.GetAllUsers()
// 	c.JSON(http.StatusOK, users)
// }
