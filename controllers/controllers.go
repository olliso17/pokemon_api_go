package controllers

import (
	"net/http"
	"pokemon_api_go/database"
	"pokemon_api_go/models"

	"github.com/gin-gonic/gin"
)

func Hellow(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, welcome to the pokemon catalog",
	})
}

// Pokemons
func CreatePokemon(c *gin.Context) {
	var pokemon models.Pokemon
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&pokemon)
	c.JSON(http.StatusOK, pokemon)
}
func ListPokemons(c *gin.Context) {
	var pokemons []models.Pokemon
	database.DB.Find(&pokemons)
	if pokemons == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "not found",
		})
	}
	c.JSON(http.StatusOK, pokemons)
}
func GetPokemonId(c *gin.Context) {
	var pokemon models.Pokemon
	id := c.Params.ByName("id")
	database.DB.First(&pokemon, id)
	if pokemon.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "there is no pokemon",
		})
		return
	}
	c.JSON(http.StatusOK, pokemon)
}

func UpdatePokemon(c *gin.Context) {
	var pokemon models.Pokemon
	id := c.Params.ByName("id")
	database.DB.First(&pokemon, id)
	if err := c.ShouldBindJSON(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Save(&pokemon)
	c.JSON(http.StatusOK, pokemon)
}

func DeletePokemon(c *gin.Context) {
	var pokemon models.Pokemon
	id := c.Params.ByName("id")
	database.DB.Delete(&pokemon, id)

}

// Types Pokemon
func CreateType(c *gin.Context) {
	var typePokemon models.Type
	if err := c.ShouldBindJSON(&typePokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&typePokemon)
	c.JSON(http.StatusOK, typePokemon)
}
func ListTypes(c *gin.Context) {
	var types []models.Type
	database.DB.Find(&types)
	if types == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "not found",
		})
	}
	c.JSON(http.StatusOK, types)
}
func GetTypeId(c *gin.Context) {
	var typePokemon models.Type
	id := c.Params.ByName("id")
	database.DB.First(&typePokemon, id)
	if typePokemon.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "there is no typePokemon",
		})
		return
	}
	c.JSON(http.StatusOK, typePokemon)
}

func UpdateType(c *gin.Context) {
	var typePokemon models.Type
	id := c.Params.ByName("id")
	database.DB.First(&typePokemon, id)
	if err := c.ShouldBindJSON(&typePokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Save(&typePokemon)
	c.JSON(http.StatusOK, typePokemon)
}

func DeleteType(c *gin.Context) {
	var typePokemon models.Type
	id := c.Params.ByName("id")
	database.DB.Delete(&typePokemon, id)

}

// create typePokemon
func CreatePokemonType(c *gin.Context) {
	var typePokemon models.PokemonsTypes
	if err := c.ShouldBindJSON(&typePokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&typePokemon)
	c.JSON(http.StatusOK, typePokemon)
}

func DeletePokemonType(c *gin.Context) {
	var typePokemon models.PokemonsTypes
	id := c.Params.ByName("id")
	database.DB.Delete(&typePokemon, id)

}

// create Images
func CreateImage(c *gin.Context) {
	var image models.Image

	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&image)
	c.JSON(http.StatusOK, image)
}

// create Video
func CreateVideo(c *gin.Context) {
	var video models.Video

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&video)
	c.JSON(http.StatusOK, video)
}
