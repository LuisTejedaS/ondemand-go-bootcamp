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
	GetById(p *model.Pokemon, id int) (*model.Pokemon, error)
	LoadPokemons() error
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

func (pI *pokemonInteractor) GetById(p *model.Pokemon, id int) (*model.Pokemon, error) {
	p, err := pI.PokemonService.FindById(p, id)
	if err != nil {
		return nil, err
	}

	return pI.PokemonPresenter.ResponsePokemon(p), nil
}

func (pI *pokemonInteractor) LoadPokemons() error {
	err := pI.PokemonService.LoadPokemons()
	if err != nil {
		return err
	}

	return nil
}
