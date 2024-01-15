package main

import (
	"github.com/OtavioValadao/api-main-go.git/database"
	"github.com/OtavioValadao/api-main-go.git/routers"
)

func main() {
	database.InitializerPostgreSQL()
	routers.Initialize()
}
