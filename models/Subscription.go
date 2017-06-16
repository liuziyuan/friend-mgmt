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
func IsSubscribes(requestorID uint, targetID uint) (isSub bool) {
	db := utils.InitDb()
	defer db.Close()
	var count int
	db.Model(&Subscription{}).Where("requestor_id = ? and target_id = ?", requestorID, targetID).Count(&count)
	if count == 0 {
		isSub = false
	} else {
		isSub = true
	}
	return
}

//IsBlocked two user is blocked
func IsBlocked(requestorID uint, targetID uint) (isBlock bool) {
	db := utils.InitDb()
	defer db.Close()
	var count int
	db.Model(&Subscription{}).Where("(requestor_id = ? and target_id = ? or requestor_id = ? and target_id = ?) and is_block = ?", requestorID, targetID, targetID, requestorID, true).Count(&count)
	if count == 0 {
		isBlock = false
	} else {
		isBlock = true
	}
	return
}

//CreateSubscription sub between requestor and target
func CreateSubscription(requestor User, target User, isBlock bool) {
	db := utils.InitDb()
	defer db.Close()
	db.Create(&Subscription{Requestor: requestor, Target: target, IsBlock: isBlock})
}

//UpdateSubscription update sub
func UpdateSubscription(sub Subscription) {
	db := utils.InitDb()
	defer db.Close()
	db.Update(&sub)
}

//GetOneSubscription ...
func GetOneSubscription(requestorID uint, targetID uint) (sub Subscription) {
	db := utils.InitDb()
	defer db.Close()
	db.Where("requestor_id = ? and target_id = ?", requestorID, targetID).First(&sub)
	return
}

//GetBlockList ...
func GetBlockList(requestorID uint) (subs []Subscription) {
	db := utils.InitDb()
	defer db.Close()
	db.Where("requestor_id = ? and is_block = ?", requestorID, true).Find(&subs)
	return
}

//GetSubscriptionListByReqID ...
func GetSubscriptionListByReqID(requestorID uint) (subs []Subscription) {
	db := utils.InitDb()
	defer db.Close()
	db.Where("requestor_id = ? and is_block = ?", requestorID, false).Find(&subs)
	return
}
