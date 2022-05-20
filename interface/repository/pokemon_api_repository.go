package repository

import (
	"encoding/json"
	"strconv"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/infraestructure/data"
)

type pokemonApiLoader struct {
	apiDataSource data.ApiDataSource
	csvDataStore  data.CsvDataStore
}

func NewPokemonAPIRepository(ds data.ApiDataSource, st data.CsvDataStore) (*pokemonApiLoader, error) {
	s := &pokemonApiLoader{apiDataSource: ds, csvDataStore: st}
	return s, nil
}

func (a *pokemonApiLoader) loadPokemons() error {
	toSave := make([][]string, 0, 5)
	records, err := a.apiDataSource.ReadCollection()
	if err != nil {
		return err
	}
	var res model.ApiResult
	json.Unmarshal([]byte(records), &res)

	for k, v := range res.Results {
		id := strconv.Itoa(k)
		name := v.Name
		// el := []string{id, name}
		toSave = append(toSave, []string{id, name})
	}

	err = a.csvDataStore.SaveRecords(toSave)
	if err != nil {
		return err
	}
	return nil
}
