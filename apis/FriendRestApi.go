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
}

//GetFriendsHandler Get All Friends
func GetFriendsHandler(c *gin.Context) {

}

//ConnectFriendsHandler create a friend connection between two email addresses.
func ConnectFriendsHandler(c *gin.Context) {
	var input dtos.FriendInput
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

		}
	}
}
