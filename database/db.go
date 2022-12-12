package database

import (
	"log"
	"os"
	"pokemon_api_go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDB() {

	dsn := os.Getenv("DBACESS")
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Pokemon{}, &models.Image{}, &models.Type{}, &models.PokemonsTypes{}, &models.Video{})
}
