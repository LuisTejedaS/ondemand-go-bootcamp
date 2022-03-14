package presenter

import "github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(p []*model.Pokemon) []*model.Pokemon {
	res := make([]model.Pokemon, len(p))
	for i, po := range p {
		res[i] = *po
	}
	return p

}
