package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var exp = "poke.csv"

func TestLoad(t *testing.T) {
	WD, err := os.Getwd()
	p := filepath.Join(WD, "../", "../")
	configuration, err := NewConfiguration(p)
	if err != nil {
		t.Errorf("Error in Pokemon service: %s", err)
	}
	c := configuration.ConfData.CSV
	if !strings.Contains(c, exp) {
		t.Errorf("Configuration reading error, got: %s, want: %s.", c, exp)
	}

}
