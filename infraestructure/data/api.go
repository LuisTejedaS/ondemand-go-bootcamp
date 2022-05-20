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
	return "", nil
}
