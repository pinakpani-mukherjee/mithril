package database

import (
	"mithril/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectWithAutoMigrate() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/mithril"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
	DB.AutoMigrate(models.User{})

}
