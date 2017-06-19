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

//GetUserByIds Get user by ids
func GetUserByIds(ids []int) (users []User) {
	db := utils.InitDb()
	defer db.Close()
	db.Where("id in (?)", ids).Find(&users)
	//db.Table("users").Select("email_address").Where("id in (?)", ids).Scan(&emails)
	return
}

//CreateUser ...
func CreateUser(email string) {
	db := utils.InitDb()
	defer db.Close()
	db.Create(&User{EmailAddress: email})
}
