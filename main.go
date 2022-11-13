package main

import (
	"log"
	"stock-manager-api/database"
	"stock-manager-api/routes"
	"stock-manager-api/utils"
)

func main() {
	err := database.CreateDbConnection()
	if err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	utils.InitMainLogger()

	app := routes.New()
	log.Fatal(app.Listen(":3001"))
}
