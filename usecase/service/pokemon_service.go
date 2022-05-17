package service

import (
	"github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/repository"
)

type pokemonService struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonService interface {
	FindAll([]*model.Pokemon) ([]*model.Pokemon, error)
	FindById(*model.Pokemon, int) (*model.Pokemon, error)
}

func NewPokemonService(r repository.PokemonRepository) pokemonService {
	return pokemonService{PokemonRepository: r}
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
