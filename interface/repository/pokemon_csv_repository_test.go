package repository

import (
	"testing"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockedCsvDataSource struct {
	mock.Mock
}

func (m MockedCsvDataSource) ReadCollection() ([][]string, error) {

	ret := m.Called()
	r0 := ret.Get(0).([][]string)
	r1 := ret.Error(1)
	return r0, r1
}

var mockCSV = [][]string{
	{"1", "Bulbasaur"},
	{"2", "Ivysaur"},
	{"3", "Venusaur"},
	{"4", "Charmander"},
}

func TestFindAll(t *testing.T) {
	var p []*model.Pokemon
	pokeDS := new(MockedCsvDataSource)
	pokeDS.On("ReadCollection").Return(mockCSV, nil)

	poCSVRepository, err := NewpokemonCSVRepository(pokeDS)
	pokemons, err := poCSVRepository.FindAll(p)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if len(pokemons) != 4 {
		t.Errorf("Pokemon reading error, got: %d, want: %d.", len(p), 4)
	}

}

func TestFindById(t *testing.T) {
	var p *model.Pokemon
	pokeDS := new(MockedCsvDataSource)
	pokeDS.On("ReadCollection").Return(mockCSV, nil)

	poCSVRepository, err := NewpokemonCSVRepository(pokeDS)
	pokemon, err := poCSVRepository.FindById(p, 4)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	if pokemon.ID != 4 {
		t.Errorf("Pokemon reading error, got: %d, want: %d.", pokemon.ID, 4)
	}

}
