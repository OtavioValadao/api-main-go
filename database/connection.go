package database

import (
	"github.com/OtavioValadao/api-main-go.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
)

func InitializerPostgreSQL() *gorm.DB {
	host := config.GetEnvironments("DB_URL")
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{})
	if err != nil {
		logger.Error("Something wrong with database connection")
		return nil
	}

	return db
}
