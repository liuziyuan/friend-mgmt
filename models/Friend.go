package models

import (
	"friends-mgmt-gin/utils"

	"github.com/jinzhu/gorm"
)

// Friend The Friend Struct
type Friend struct {
	gorm.Model
	Requestor   User
	RequestorID int

	Target   User
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
	db.Create(&Friend{Requestor: requestor, Target: target})
}

//RetrieveFriends retrieve the friends list for an email address
func RetrieveFriends(uid uint) (friends []Friend) {
	db := utils.InitDb()
	defer db.Close()
	db.Where("requestor_id = ?", uid).Or("target_id = ?", uid).Find(&friends)
	return
}

//GetFriendsByEmail Get Friends by Email
func GetFriendsByEmail(email string) (uids []int) {
	user := GetUserByEmailAddr(email)
	friends := RetrieveFriends(user.ID)
	alice := []int{}
	for _, friend := range friends {
		if friend.RequestorID == int(user.ID) {
			alice = append(alice, friend.TargetID)
		} else if friend.TargetID == int(user.ID) {
			alice = append(alice, friend.RequestorID)
		}
	}
	return alice
}

//GetFriendsEmail ...
func GetEmails(users []User) (emails []string) {
	for _, user := range users {
		emails = append(emails, user.EmailAddress)
	}
	return
}
