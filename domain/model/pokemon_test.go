package model

import (
	"testing"
)

func TestCreatePokemonConstructor(t *testing.T) {
	pID := uint(1)
	pName := "bulbasaur"
	cPokemon := newPokemon(uint(pID), pName)
	if cPokemon.ID != pID {
		t.Errorf("Pokemon Id was not set, got: %d, want: %d.", cPokemon.ID, pID)
	}

	if cPokemon.Name != pName {
		t.Errorf("Pokemon Name was not set, got: %s, want: %s.", cPokemon.Name, pName)
	}
}

func TestCreatePokemon(t *testing.T) {
	pID := uint(1)
	pName := "bulbasaur"
	cPokemon := Pokemon{uint(pID), pName}
	if cPokemon.ID != pID {
		t.Errorf("Pokemon Id was not set, got: %d, want: %d.", cPokemon.ID, pID)
	}

	if cPokemon.Name != pName {
		t.Errorf("Pokemon Name was not set, got: %s, want: %s.", cPokemon.Name, pName)
	}
}
