package data

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type apiDataSource struct {
	baseUrl string
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
	url := c.baseUrl + "?offset=0&limit=151"

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}
