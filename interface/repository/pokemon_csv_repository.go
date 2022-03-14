package repository

import (
	"errors"
	"strconv"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/infraestructure/data"
)

type pokemonCSVRepository struct {
	csvDataSource data.CsvDataSource
	pokemons      []*model.Pokemon
}

func NewpokemonCSVRepository(ds data.CsvDataSource) (*pokemonCSVRepository, error) {
	adapter := &pokemonCSVRepository{csvDataSource: ds}

	if err := adapter.loadPokemons(); err != nil {
		return nil, err
	}

	return adapter, nil
}

func (a *pokemonCSVRepository) loadPokemons() error {
	csvRecords, err := a.csvDataSource.ReadCollection()
	if err != nil {
		return err
	}

	for _, v := range csvRecords {
		p := model.Pokemon{}

		strconv.Atoi(v[0])
		p.Name = v[1]

		a.pokemons = append(a.pokemons, &p)
	}

	return nil
}

// GetByID searches for a pokemon with the given id parameter.
//
// If the search is successful, a pointer to the found Pokemon is returned.
// Otherwise and ErrPokemonNotFoundByID error is returned.
// func (a *pokemonCSVRepository) GetByID(id int) (*entity.Pokemon, error) {
// 	for _, pokemon := range a.pokemons {
// 		if id == int(pokemon.ID) {
// 			return &pokemon, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("%w %v", ErrPokemonNotFoundByID, id)
// }

// GetAll returns a slice of all pokemons.
//
// In case no pokemons are found at all, an ErrPokemonsNotFound error is returned.
func (a *pokemonCSVRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	if len(a.pokemons) == 0 {
		return nil, errors.New("no pokemons found")
	}

	return a.pokemons, nil
}
