package config

import (
	"testing"
)

var exp = "/Users/luis.tejeda/Documents/Source/GitHub/ondemand-go-bootcamp/poke.csv"

func TestLoad(t *testing.T) {
	configuration, err := NewConfiguration("/Users/luis.tejeda/Documents/Source/GitHub/ondemand-go-bootcamp/")
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	c := configuration.ConfData.Csv

	if c != exp {
		t.Errorf("Configuration reading error, got: %s, want: %s.", c, exp)
	}

}
