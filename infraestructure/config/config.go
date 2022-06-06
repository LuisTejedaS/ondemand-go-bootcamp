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
	WD       string
}

type confData struct {
	CSV         string `json:"csv"`
	PokeBaseURL string `json:"pokeBaseUrl"`
	Port        string `json:"port"`
}

func NewConfiguration(WD string) (configuration, error) {
	var conf = configuration{WD: WD}
	err := conf.Load()
	if err != nil {
		return conf, err
	}
	return conf, nil
}

func (c *configuration) Load() error {
	jsonFile, err := os.Open(c.WD + "/conf.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var d confData
	err = json.Unmarshal(byteValue, &d)
	if err != nil {
		return err
	}

	c.ConfData = d
	return nil
}
