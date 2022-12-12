package main

import (
	"pokemon_api_go/database"
	"pokemon_api_go/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.ConectDB()
	routes.HandleRequest()
}
