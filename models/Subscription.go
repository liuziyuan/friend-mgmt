package models

import (
	"friends-mgmt-gin/utils"

	"github.com/jinzhu/gorm"
)

//Subscription ...
type Subscription struct {
	gorm.Model

	Requestor   User
	RequestorID int

	Target   User
	TargetID int

	IsBlock bool
}

//IsSubscribes The subscribe relationship exited
func IsSubscribes(requestorID uint, targetID uint) (isFriends bool) {
	db := utils.InitDb()
	defer db.Close()
	var count int
	db.Model(&Subscription{}).Where("requestor_id = ? and target_id = ?", requestorID, targetID).Count(&count)
	if count == 0 {
		isFriends = false
	} else {
		isFriends = true
	}
	return
}

//CreateSubscription sub between requestor and target
func CreateSubscription(requestor User, target User) {
	db := utils.InitDb()
	defer db.Close()
	db.Create(&Subscription{Requestor: requestor, Target: target, IsBlock: false})
}
