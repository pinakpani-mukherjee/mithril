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
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       rand.Float64(),
		}
		product.Id = uint(rand.Intn(max-min+1) + min)

		database.DB.Create(&product)
	}
}
