package repository

import "github.com/LuisTejedaS/ondemand-go-bootcamp/domain/model"

type PokemonRepository interface {
	FindAll(u []*model.Pokemon) ([]*model.Pokemon, error)
	//FindById(u []*model.Pokemon) (*model.Pokemon, error)
}
