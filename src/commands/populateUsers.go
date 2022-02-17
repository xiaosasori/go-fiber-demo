package main

import (
	"go-fiber-demo/src/database"
	"go-fiber-demo/src/models"

	"github.com/bxcodec/faker/v3"
)

// this must run inside docker container
func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}
}
