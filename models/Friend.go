package models

import (
	"friends-mgmt-gin/utils"

	"github.com/jinzhu/gorm"
)

// Friend The Friend Struct
type Friend struct {
	gorm.Model
	Requestor   *User
	RequestorID int

	Target   *User
	TargetID int
}

//IsFriends The friend relationship exited
func IsFriends(requestorID uint, targetID uint) (isFriends bool) {
	db := utils.InitDb()
	defer db.Close()
	var count int
	db.Model(&Friend{}).Where("requestor_id = ? and target_id = ?", requestorID, targetID).Or("target_id = ? and requestor_id = ?", requestorID, targetID).Count(&count)
	if count == 0 {
		isFriends = false
	} else {
		isFriends = true
	}
	return
}

//CreateFriends Connect user each others
func CreateFriends(requestor User, target User) {
	db := utils.InitDb()
	defer db.Close()
	db.Create(&Friend{Requestor: &requestor, Target: &target})
}