package main

import (
	"github.com/OtavioValadao/api-main-go.git/database"
	"github.com/OtavioValadao/api-main-go.git/models"
)

func main() {
	database.InitializerPostgreSQL().AutoMigrate(
		&models.AnimalGuardian{},
		&models.Animal{},
	)
}
