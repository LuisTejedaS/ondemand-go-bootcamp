package service

import (
	"testing"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockedRepository struct {
	mock.Mock
}

func (m MockedRepository) FindAll(u []*model.Pokemon) ([]*model.Pokemon, error) {
	ret := m.Called()

	r0 := ret.Get(0).([]*model.Pokemon)
	r1 := ret.Error(1)

	return r0, r1
}

func (m MockedRepository) FindById(u *model.Pokemon, id int) (*model.Pokemon, error) {
	ret := m.Called()

	r0 := ret.Get(0).(*model.Pokemon)
	r1 := ret.Error(1)

	return r0, r1
}

var mockPokemon = []*model.Pokemon{
	{ID: 1, Name: "Bulbasaur"},
	{ID: 2, Name: "Ivysaur"},
	{ID: 3, Name: "Venusaur"},
	{ID: 4, Name: "Charmander"},
}

func TestFindAll(t *testing.T) {
	var p []*model.Pokemon
	pokeRepository := new(MockedRepository)
	pokeRepository.On("FindAll").Return(mockPokemon, nil)

	pokemonService := NewPokemonService(pokeRepository)
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
	pokeRepository := new(MockedRepository)
	pokeRepository.On("FindById").Return(mockPokemon[2], nil)

	pokemonService := NewPokemonService(pokeRepository)
	p, err := pokemonService.FindById(p, 3)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if p.ID != 3 {
		t.Errorf("Pokemon reading error id, got: %d, want: %d.", p.ID, 4)
	}

}
