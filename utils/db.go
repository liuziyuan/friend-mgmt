package utils

import (
	"github.com/jinzhu/gorm"
)

// InitDb Init Db Connention
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("postgres", "host=192.168.10.199 user=postgres password=1qaz2wsx3EDC dbname=Friend_Gin sslmode=disable ")
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database.")
	}
	return db
}
