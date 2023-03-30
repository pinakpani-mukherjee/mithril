package main

import (
	"math/rand"
	"mithril/src/database"
	"mithril/src/models"

	"github.com/go-faker/faker/v4"
)

func main() {
	database.ConnectWithAutoMigrate()
	min := 100000000
	max := 999999999
	for i := 0; i < 30; i++ {
		mithril := models.User{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			IsMithril: true,
		}
		mithril.Id = uint(rand.Intn(max-min+1) + min)

		mithril.SetPassword("123456")

		database.DB.Create(&mithril)
	}
}
