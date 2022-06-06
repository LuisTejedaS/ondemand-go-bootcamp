package model

type ApiResult struct {
	Count    int64        `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []PokeResult `json:"results"`
}

type PokeResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
