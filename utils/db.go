package utils

import (
	"github.com/jinzhu/gorm"
)

// InitDb Init Db Connention
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("postgres", "host=192.168.10.199 user=postgres password=1qaz2wsx3EDC dbname=Friend_Gin sslmode=disable ")

	// db, err := gorm.Open("postgres", "host=10.120.8.137 user=u8ae6c2d62d74477a84dfa6f693a45d3d password=4a76ff5639aa49a3a00d95bbd4119857 dbname=d9488550a0c6e4feabc9a1f8c6db569d6 sslmode=disable ")

	db.LogMode(true)
	if err != nil {
		panic("failed to connect database.")
	}
	return db
}
