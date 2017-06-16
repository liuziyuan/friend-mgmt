package restapis

import (
	"friends-mgmt-gin/dtos"
	"friends-mgmt-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Callback ...
type Callback func(requestor models.User, target models.User, c *gin.Context)

var (
	// Subscriptions Subscription Router Group
	Subscriptions *gin.RouterGroup
)

//AddSubscriptionRoutes ...
func AddSubscriptionRoutes() {
	Subscriptions.GET("/", GetSubscriptionsHandler)
	Subscriptions.POST("/subscribe", SubscribeHandler)
	Subscriptions.POST("/block", BlockSubscribeHandler)
	Subscriptions.POST("/retrieve", RetrieveSubscribeHandler)
}

//GetSubscriptionsHandler ...
func GetSubscriptionsHandler(c *gin.Context) {

}

//SubscribeHandler subscribe to updates from an email address
func SubscribeHandler(c *gin.Context) {
	CommonSubscribeHandler(c, SubscribeLogic)
}

//SubscribeLogic private method
func SubscribeLogic(requestor models.User, target models.User, c *gin.Context) {
	isSub := models.IsSubscribes(requestor.ID, target.ID)
	if !isSub {
		models.CreateSubscription(requestor, target, false)
		c.JSON(http.StatusOK, gin.H{"success": "true"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "already subscribed"})
	}
}

//BlockSubscribeHandler block updates from an email address.
// Suppose "andy@example.com" blocks "john@example.com":
// • if they are connected as friends, then "andy" will no longer receive notifications from
// "john"
// • if they are not connected as friends, then no new friends connection can be added
func BlockSubscribeHandler(c *gin.Context) {
	CommonSubscribeHandler(c, BlockSubscribeLogic)
}

//BlockSubscribeLogic private method
func BlockSubscribeLogic(requestor models.User, target models.User, c *gin.Context) {
	isSub := models.IsSubscribes(requestor.ID, target.ID)
	if !isSub {
		models.CreateSubscription(requestor, target, true)
	} else {
		sub := models.GetOneSubscription(requestor.ID, target.ID)
		sub.IsBlock = true
		models.UpdateSubscription(sub)
	}
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

//RetrieveSubscribeHandler retrieve all email addresses that can receive updates from an email address
// Eligibility for receiving updates from i.e. "john@example.com":
// • has not blocked updates from "john@example.com", and
// • at least one of the following:
// o has a friend connection with "john@example.com"
// o has subscribed to updates from "john@example.com"
// o has been @mentioned in the update
func RetrieveSubscribeHandler(c *gin.Context) {

}

//CommonSubscribeHandler The common function ,you need to call your logic function by callback
func CommonSubscribeHandler(c *gin.Context, callback Callback) {
	var input dtos.SubscriptionInput
	if c.BindJSON(&input) == nil {
		requestor := models.GetUserByEmailAddr(input.Requestor)
		target := models.GetUserByEmailAddr(input.Target)
		if &requestor != nil && &target != nil {
			callback(requestor, target, c)
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "someone is not system user"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please make sure the parms is right"})
	}
}
