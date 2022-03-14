package router

import (
	"net/http"

	controller "github.com/LuisTejedaS/ondemand-go-bootcamp/interface/controllers"
	"github.com/gin-gonic/gin"
)

type appRouter struct {
	controller controller.PokemonController
}

type AppRouter interface {
	Router() http.Handler
}

// NewAppRouter returns a new instance of AppRouter
func NewRoute(c controller.PokemonController) AppRouter {
	return &appRouter{c}
}

// Router returns a httprouter.Router with all routes configured
func (c *appRouter) Router() http.Handler {
	router := gin.Default()
	router.GET("/Pokemons", c.controller.GetPokemons)
	return router
}
