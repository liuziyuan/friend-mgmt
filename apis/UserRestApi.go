package restapis

import (
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

func PostUserHandler(c *gin.Context) {

}
