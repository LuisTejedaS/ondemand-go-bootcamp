package data

import (
	"errors"
)

type apiDataSource struct {
	csvPath string
}

type ApiDataSource interface {
	ReadCollection() (string, error)
}

func NewApiDataSource(url string) (apiDataSource, error) {
	if url == "" {
		return apiDataSource{url}, errors.New("No url sent")
	}
	return apiDataSource{url}, nil
}

func (c apiDataSource) ReadCollection() (string, error) {
	return `{
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
	}`, nil
}
