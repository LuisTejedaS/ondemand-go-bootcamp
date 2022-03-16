package presenter

import "github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(p []*model.Pokemon) []*model.Pokemon {
	return p
}

func (pp *pokemonPresenter) ResponsePokemon(p *model.Pokemon) *model.Pokemon {
	return p
}
