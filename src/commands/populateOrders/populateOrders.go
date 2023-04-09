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
		var orderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5))

			OrderItem := models.OrderItem{
				ProductTitle:   faker.Word(),
				Price:          price,
				Quantity:       qty,
				AdminRevenue:   0.9 * price * float64(qty),
				MithrilRevenue: 0.1 * price * float64(qty),
			}
			OrderItem.Id = uint(rand.Intn(max-min+1) + min)

			orderItems = append(orderItems, OrderItem)
		}
		Order := models.Order{
			UserId:       uint(rand.Intn(max-min+1) + min),
			Code:         faker.Username(),
			MithrilEmail: faker.Email(),
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			Complete:     true,
			OrderItems:   orderItems,
		}
		Order.Id = uint(rand.Intn(max-min+1) + min)
		database.DB.Create(&Order)

	}
}
