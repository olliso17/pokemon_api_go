package routes

import (
	"pokemon_api_go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", controllers.Hellow)
	// routes Pokemon
	r.POST("/pokemons/create", controllers.CreatePokemon)
	r.GET("/pokemons", controllers.ListPokemons)
	r.GET("/pokemons/:id", controllers.GetPokemonId)
	r.PUT("/pokemons/:id", controllers.UpdatePokemon)
	r.DELETE("/pokemons/:id", controllers.DeletePokemon)

	// routes Skill
	r.POST("/type/create", controllers.CreateType)
	r.GET("/types", controllers.ListTypes)
	r.GET("/types/:id", controllers.GetTypeId)
	r.PUT("/types/:id", controllers.UpdateType)
	r.DELETE("/types/:id", controllers.DeleteType)

	//routes PokemonsTypes
	r.POST("/pokemontype/create", controllers.CreatePokemonType)
	r.DELETE("/pokemontype/:id", controllers.DeletePokemonType)

	// routes Images
	r.POST("/images/create", controllers.CreateImage)

	r.Run()
}
