package main

import (
	"fmt"
	"friends-mgmt-gin/apis"
	"friends-mgmt-gin/models"
	"friends-mgmt-gin/utils"
	"os"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	g := gin.Default()
	g.Use(corsHandler)
	gin.SetMode(gin.ReleaseMode)
	g.StaticFS("swagger", http.Dir("swagger"))
	restapis.Users = g.Group("/api/users")
	restapis.Friends = g.Group("/api/friends")
	restapis.Subscriptions = g.Group("/api/subscriptions")
	restapis.AddUserRoutes()
	restapis.AddFriendRoutes()
	restapis.AddSubscriptionRoutes()

	db := utils.InitDb()
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Friend{}, &models.Subscription{})

	if os.Getenv("PORT") == "" {
		g.Run(":8080")
	} else {
		g.Run(":" + os.Getenv("PORT"))
	}

}

func corsHandler(c *gin.Context) {
	fmt.Println("RequestURI :")
	if strings.HasPrefix(c.Request.RequestURI, "/user") {
		c.Header("Content-Type", "application/json")
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept, authorization, Pragma, Cache-control, Expires")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	c.Header("Access-Control-Max-Age", "1209600")
}
