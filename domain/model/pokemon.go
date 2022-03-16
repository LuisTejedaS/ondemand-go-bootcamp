package model

type Pokemon struct {
	ID   uint64
	Name string
}

func newPokemon(Id uint64, name string) *Pokemon {
	p := Pokemon{Name: name, ID: Id}
	return &p
}
