package repository

import (
	"testing"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	mocks "github.com/LuisTejedaS/ondemand-go-bootcamp/mocks/infraestructure/data"
)

var mockCSV = [][]string{
	{"1", "Bulbasaur"},
	{"2", "Ivysaur"},
	{"3", "Venusaur"},
	{"4", "Charmander"},
}

var mockCSVEmpty = [][]string{}

func TestFindAll(t *testing.T) {
	var p []*model.Pokemon
	pokeDS := new(mocks.CsvDataSource)
	pokeDS.On("ReadCollection").Return(mockCSV, nil)

	poCSVRepository, err := NewpokemonCSVRepository(pokeDS)
	pokemons, err := poCSVRepository.FindAll(p)
	if err != nil {
		t.Errorf("Error in Pokemon csv repository: %s", err)
	}
	if len(pokemons) != 4 {
		t.Errorf("Pokemon reading error, got: %d, want: %d.", len(p), 4)
	}

}

func TestFindById(t *testing.T) {
	var p *model.Pokemon
	pokeDS := new(mocks.CsvDataSource)
	pokeDS.On("ReadCollection").Return(mockCSV, nil)

	poCSVRepository, err := NewpokemonCSVRepository(pokeDS)
	pokemon, err := poCSVRepository.FindById(p, 4)
	if err != nil {
		t.Errorf("Error in Pokemon csv repository: %s", err)
	}
	if pokemon.ID != 4 {
		t.Errorf("Pokemon reading error, got: %d, want: %d.", pokemon.ID, 4)
	}

}

func TestFindByIdNotFound(t *testing.T) {
	var p *model.Pokemon
	pokeDS := new(mocks.CsvDataSource)
	pokeDS.On("ReadCollection").Return(mockCSV, nil)

	poCSVRepository, err := NewpokemonCSVRepository(pokeDS)
	pokemon, err := poCSVRepository.FindById(p, 5)
	if err.Error() != "no pokemon found 5" {
		t.Errorf("Error in Pokemon csv repository: %s", err)
	}
	if pokemon != nil {
		t.Errorf("Error in Pokemon csv repository: should not find any and found %d", pokemon.ID)
	}

}

func TestFindAllNotFound(t *testing.T) {
	var p []*model.Pokemon
	pokeDS := new(mocks.CsvDataSource)
	pokeDS.On("ReadCollection").Return(mockCSVEmpty, nil)

	poCSVRepository, err := NewpokemonCSVRepository(pokeDS)
	pokemons, err := poCSVRepository.FindAll(p)
	if err.Error() != "no pokemons found" {
		t.Errorf("Error in Pokemon csv repository: %s", err)
	}
	if pokemons != nil {
		t.Errorf("Error in Pokemon csv repository: should not find any and found %d", len(pokemons))
	}

}
