package repository

type PokemonLoader interface {
	LoadPokemons() error
}
