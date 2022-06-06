package repository

import (
	"testing"

	mocks "github.com/LuisTejedaS/ondemand-go-bootcamp/mocks/infraestructure/data"
)

var apiRes = `{
	"count": 1126,
	"next": "https://pokeapi.co/api/v2/pokemon?offset=150&limit=150",
	"previous": null,
	"results": [
			{
					"name": "bulbasaur",
					"url": "https://pokeapi.co/api/v2/pokemon/1/"
			},
			{
				"name": "ivysaur",
				"url": "https://pokeapi.co/api/v2/pokemon/2/"
		}  ]
}`
var srEx = [][]string{{"id", "name"}, {"0", "bulbasaur"}, {"1", "ivysaur"}}

func TestLoadPokemons(t *testing.T) {
	pokeDS := new(mocks.ApiDataSource)
	pokeStore := new(mocks.CsvDataStore)
	pokeDS.On("ReadCollection").Return(apiRes, nil)
	pokeStore.On("DeleteRecords").Return(nil)
	pokeStore.On("SaveRecords", srEx).Return(nil)

	poApiRepository, err := NewPokemonApiLoader(pokeDS, pokeStore)
	if err != nil {
		t.Errorf("Error creating api repository: %s", err)
	}
	err = poApiRepository.LoadPokemons()
	if err != nil {
		t.Errorf("Error loading pokemon: %s", err)
	}
}
