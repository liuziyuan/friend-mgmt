package restapis

import (
	"friends-mgmt-gin/dtos"
	"friends-mgmt-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// Friends Friend Router Group
	Friends *gin.RouterGroup
)

//AddFriendRoutes Add Friends RestApi Routes
func AddFriendRoutes() {
	Friends.GET("/", GetFriendsHandler)
	Friends.POST("/connect", ConnectFriendsHandler)
	Friends.POST("/retrieve", RetrieveFriendsHandler)
}

//GetFriendsHandler Get All Friends
func GetFriendsHandler(c *gin.Context) {

}

//ConnectFriendsHandler create a friend connection between two email addresses.
func ConnectFriendsHandler(c *gin.Context) {
	var input dtos.FriendsInput
	if c.BindJSON(&input) == nil {
		requestor := models.GetUserByEmailAddr(input.Friends[0])
		target := models.GetUserByEmailAddr(input.Friends[1])

		if &requestor != nil && &target != nil {
			isFriends := models.IsFriends(requestor.ID, target.ID)
			if !isFriends {
				models.CreateFriends(requestor, target)
				c.JSON(http.StatusOK, gin.H{"success": "true"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "already friends"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "someone is not system user"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please make sure the parms is right"})
	}
}

//RetrieveFriendsHandler retrieve the friends list for an email address
func RetrieveFriendsHandler(c *gin.Context) {
	var input dtos.RetrieveInput
	if c.BindJSON(&input) == nil {
		//search
		user := models.GetUserByEmailAddr(input.Email)
		friends := models.RetrieveFriends(user.ID)
		alice := []int{}
		for _, friend := range friends {
			if friend.RequestorID == int(user.ID) {
				alice = append(alice, friend.TargetID)
			} else if friend.TargetID == int(user.ID) {
				alice = append(alice, friend.RequestorID)
			}
		}
		users := models.GetUserByIds(alice)
		emails := []string{}
		for _, user := range users {
			emails = append(emails, user.EmailAddress)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
			"friends": emails,
			"count":   len(alice),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please make sure the parms is right"})
	}
}
