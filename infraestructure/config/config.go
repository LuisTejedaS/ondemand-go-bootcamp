package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configuration interface {
	Load() error
}

type configuration struct {
	ConfData confData
	Wd       string
}

type confData struct {
	Csv         string `json:"csv"`
	PokeBaseUrl string `json:"pokeBaseUrl"`
}

func NewConfiguration(wd string) (configuration, error) {
	var conf = configuration{Wd: wd}
	err := conf.Load()
	if err != nil {
		return conf, err
	}
	return conf, nil
}

func (c *configuration) Load() error {
	jsonFile, err := os.Open(c.Wd + "/conf.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var d confData
	json.Unmarshal(byteValue, &d)

	c.ConfData = d
	return nil

}
