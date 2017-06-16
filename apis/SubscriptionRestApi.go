package restapis

import (
	"friends-mgmt-gin/dtos"
	"friends-mgmt-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// Subscriptions Subscription Router Group
	Subscriptions *gin.RouterGroup
)

//AddSubscriptionRoutes ...
func AddSubscriptionRoutes() {
	Subscriptions.GET("/", GetSubscriptionsHandler)
	Subscriptions.POST("/subscribe", SubscribeHandler)
	// Subscriptions.POST("/retrieve", RetrieveFriendsHandler)
	// Subscriptions.POST("/common", RetrieveCommonFriendsHandler)
}

//GetSubscriptionsHandler ...
func GetSubscriptionsHandler(c *gin.Context) {

}

//SubscribeHandler subscribe to updates from an email address
func SubscribeHandler(c *gin.Context) {
	var input dtos.SubscriptionInput
	if c.BindJSON(&input) == nil {
		requestor := models.GetUserByEmailAddr(input.Requestor)
		target := models.GetUserByEmailAddr(input.Target)

		if &requestor != nil && &target != nil {
			isSub := models.IsSubscribes(requestor.ID, target.ID)
			if !isSub {
				models.CreateSubscription(requestor, target)
				c.JSON(http.StatusOK, gin.H{"success": "true"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "already subscribed"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "someone is not system user"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please make sure the parms is right"})
	}
}
