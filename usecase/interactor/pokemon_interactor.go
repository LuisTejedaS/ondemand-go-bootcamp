package interactor

import (
	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/presenter"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/service"
)

type pokemonInteractor struct {
	PokemonService   service.PokemonService
	PokemonPresenter presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(s service.PokemonService, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{s, p}
}

func (pI *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := pI.PokemonService.FindAll(p)
	if err != nil {
		return nil, err
	}

	return pI.PokemonPresenter.ResponsePokemons(p), nil
}
