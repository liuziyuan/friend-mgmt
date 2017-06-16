package main

import (
	"friends-mgmt-gin/apis"
	"friends-mgmt-gin/models"
	"friends-mgmt-gin/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	g := gin.Default()

	g.Static("swagger", "E:/GoWorks/src/friends-mgmt-gin/swagger")
	restapis.Users = g.Group("/api/users")
	restapis.Friends = g.Group("/api/friends")
	restapis.AddUserRoutes()
	restapis.AddFriendRoutes()

	db := utils.InitDb()
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Friend{})

	// db.Create(&models.User{EmailAddress: "mark"})
	// db.Create(&models.User{EmailAddress: "lee"})
	// db.Create(&models.User{EmailAddress: "angela"})
	// db.Create(&models.User{EmailAddress: "jason"})
	// db.Create(&models.User{EmailAddress: "eric"})
	g.Run(":8080")
}
