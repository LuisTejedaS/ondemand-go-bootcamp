package service

import (
	"testing"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	mocks "github.com/LuisTejedaS/ondemand-go-bootcamp/mocks/usecase/repository"
)

var mockPokemon = []*model.Pokemon{
	{ID: 1, Name: "Bulbasaur"},
	{ID: 2, Name: "Ivysaur"},
	{ID: 3, Name: "Venusaur"},
	{ID: 4, Name: "Charmander"},
}

func TestFindAll(t *testing.T) {
	var p []*model.Pokemon
	pokeRepository := new(mocks.PokemonRepository)
	pokeLoader := new(mocks.PokemonLoader)
	pokeRepository.On("FindAll", p).Return(mockPokemon, nil)

	pokemonService := NewPokemonService(pokeRepository, pokeLoader)
	p, err := pokemonService.FindAll(p)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if len(p) != 4 {
		t.Errorf("Pokemon reading error, got: %d, want: %d.", len(p), 4)
	}

}

func TestFindByID(t *testing.T) {
	var p *model.Pokemon
	pokeRepository := new(mocks.PokemonRepository)
	pokeLoader := new(mocks.PokemonLoader)

	pokeRepository.On("FindById", p, 3).Return(mockPokemon[2], nil)

	pokemonService := NewPokemonService(pokeRepository, pokeLoader)
	p, err := pokemonService.FindById(p, 3)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if p.ID != 3 {
		t.Errorf("Pokemon reading error id, got: %d, want: %d.", p.ID, 4)
	}

}
