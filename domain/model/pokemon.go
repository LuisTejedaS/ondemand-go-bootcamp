package model

type Pokemon struct {
	ID   uint
	Name string
}

func newPokemon(Id uint, name string) *Pokemon {
	p := Pokemon{Name: name, ID: Id}
	return &p
}
