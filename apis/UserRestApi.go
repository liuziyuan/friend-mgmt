package restapis

import (
	"friends-mgmt-gin/dtos"
	"friends-mgmt-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// Users User Router Group
	Users *gin.RouterGroup
)

// AddUserRoutes Add User Restapi Routes
func AddUserRoutes() {
	Users.GET("/", GetUsersHandler)
	Users.POST("/", PostUserHandler)
}

// GetUsersHandler Get All Users
func GetUsersHandler(c *gin.Context) {

	c.String(http.StatusOK, "Success")
}

// PostUserHandler Create a user
func PostUserHandler(c *gin.Context) {
	var input dtos.RetrieveInput
	if c.BindJSON(&input) == nil {
		user := models.GetUserByEmailAddr(input.Email)
		if user.ID == 0 {
			models.CreateUser(input.Email)
			c.JSON(http.StatusOK, gin.H{"success": "true"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "already exist"})
		}
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "please make sure the parms is right"})
	}
}
