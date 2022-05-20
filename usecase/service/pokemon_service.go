package service

import (
	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/repository"
)

type pokemonService struct {
	PokemonRepository repository.PokemonRepository
	PokemonLoader     repository.PokemonLoader
}

type PokemonService interface {
	FindAll([]*model.Pokemon) ([]*model.Pokemon, error)
	FindById(*model.Pokemon, int) (*model.Pokemon, error)
	LoadPokemons() error
}

func NewPokemonService(r repository.PokemonRepository, l repository.PokemonLoader) pokemonService {
	return pokemonService{PokemonRepository: r, PokemonLoader: l}
}

func (s pokemonService) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := s.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s pokemonService) FindById(p *model.Pokemon, id int) (*model.Pokemon, error) {
	p, err := s.PokemonRepository.FindById(p, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s pokemonService) LoadPokemons() error {
	err := s.PokemonLoader.LoadPokemons()
	if err != nil {
		return err
	}
	err = s.PokemonRepository.LoadPokemons()
	if err != nil {
		return err
	}
	return nil
}
