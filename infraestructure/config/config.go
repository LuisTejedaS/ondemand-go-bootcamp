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
	confData
}

type confData struct {
	Csv string `json:"csv"`
}

func NewConfiguration() (configuration, error) {
	var conf = configuration{}
	err := conf.Load()
	if err != nil {
		return conf, err
	}
	return conf, nil
}

func (c *configuration) Load() error {
	jsonFile, err := os.Open("/Users/luis.tejeda/Documents/Source/GitHub/ondemand-go-bootcamp/conf.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var d confData
	json.Unmarshal(byteValue, &d)

	c.confData = d
	return nil

}
