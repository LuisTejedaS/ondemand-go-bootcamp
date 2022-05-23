package interactor

import (
	"testing"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	mocksp "github.com/LuisTejedaS/ondemand-go-bootcamp/mocks/usecase/presenter"
	mocks "github.com/LuisTejedaS/ondemand-go-bootcamp/mocks/usecase/service"
)

var mockPokemon = []*model.Pokemon{
	{ID: 1, Name: "Bulbasaur"},
	{ID: 2, Name: "Ivysaur"},
	{ID: 3, Name: "Venusaur"},
	{ID: 4, Name: "Charmander"},
}

func TestGet(t *testing.T) {
	var p []*model.Pokemon
	pokeService := new(mocks.PokemonService)
	pokePresenter := new(mocksp.PokemonPresenter)

	pokeService.On("FindAll", p).Return(mockPokemon, nil)
	pokePresenter.On("ResponsePokemons", mockPokemon).Return(mockPokemon, nil)

	pokemonInteractor := NewPokemonInteractor(pokeService, pokePresenter)
	p, err := pokemonInteractor.Get(p)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if len(p) != 4 {
		t.Errorf("Pokemon reading error, got: %d, want: %d.", len(p), 4)
	}

}

func TestGetById(t *testing.T) {
	var p *model.Pokemon
	pokeService := new(mocks.PokemonService)
	pokePresenter := new(mocksp.PokemonPresenter)

	pokeService.On("FindById", p, 3).Return(mockPokemon[2], nil)
	pokePresenter.On("ResponsePokemon", mockPokemon[2]).Return(mockPokemon[2], nil)

	pokemonInteractor := NewPokemonInteractor(pokeService, pokePresenter)
	p, err := pokemonInteractor.GetById(p, 3)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if p.ID != 3 {
		t.Errorf("Pokemon reading error id, got: %d, want: %d.", p.ID, 4)
	}

}

func TestGetLoadPokemons(t *testing.T) {
	pokeService := new(mocks.PokemonService)
	pokePresenter := new(mocksp.PokemonPresenter)

	pokeService.On("LoadPokemons").Return(nil)

	pokemonInteractor := NewPokemonInteractor(pokeService, pokePresenter)
	err := pokemonInteractor.LoadPokemons()
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
}
