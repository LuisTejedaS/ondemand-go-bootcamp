package router

import (
	controller "github.com/LuisTejedaS/ondemand-go-bootcamp/interface/controllers"
	"github.com/gin-gonic/gin"
)

func getAllPokemons(g *gin.RouterGroup, p controller.PokemonController) {
	g.GET("/pokemons", p.GetPokemons)
}

func RegisterPokemonRoutes(g *gin.RouterGroup, p controller.PokemonController) {
	getAllPokemons(g, p)
}
