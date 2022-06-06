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

func NewPokemonApiLoader(ds data.ApiDataSource, st data.CsvDataStore) (*pokemonApiLoader, error) {
	s := &pokemonApiLoader{apiDataSource: ds, csvDataStore: st}
	return s, nil
}

func (a *pokemonApiLoader) LoadPokemons() error {
	toSave := make([][]string, 0, 5)
	toSave = append(toSave, []string{"id", "name"})

	records, err := a.apiDataSource.ReadCollection()
	if err != nil {
		return err
	}
	var res model.ApiResult
	json.Unmarshal([]byte(records), &res)

	for k, v := range res.Results {
		id := strconv.Itoa(k)
		name := v.Name
		toSave = append(toSave, []string{id, name})
	}

	err = a.csvDataStore.DeleteRecords()
	if err != nil {
		return err
	}
	err = a.csvDataStore.SaveRecords(toSave)
	if err != nil {
		return err
	}
	return nil
}
