package controller

import (
	"net/http"
	"strconv"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/interactor"
	"github.com/gin-gonic/gin"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c *gin.Context)
	GetPokemon(c *gin.Context)
}

func NewPokemonController(pI interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pI}
}

func (pc *pokemonController) GetPokemons(c *gin.Context) {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
	}

	c.JSON(http.StatusOK, p)
	return
}

func (pc *pokemonController) GetPokemon(c *gin.Context) {
	var p *model.Pokemon

	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	p, err = pc.pokemonInteractor.GetById(p, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
	return
}
