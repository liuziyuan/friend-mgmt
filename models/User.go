package models

import (
	"friends-mgmt-gin/utils"

	"github.com/jinzhu/gorm"
)

// User The User Struct
type User struct {
	gorm.Model
	EmailAddress string `gorm:"type:varchar(100)"`
}

//GetUserByEmailAddr Get user by Email address
func GetUserByEmailAddr(emailAddr string) (user User) {
	db := utils.InitDb()
	defer db.Close()
	//db.Debug().First(&user, "email_address = ?", emailAddr)
	db.Find(&user, "email_address = ?", emailAddr).First(&user)
	return
}
