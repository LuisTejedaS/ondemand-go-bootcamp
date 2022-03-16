package data

import (
	"testing"
)

func TestCreateNewCSVDataSource(t *testing.T) {
	p := "ondemand-go-bootcamp/poke.csv"
	csvDS, _ := NewCSVDataSource(p)
	if csvDS.csvPath != p {
		t.Errorf("CSVPath was not set, got: %s, want: %s.", csvDS.csvPath, p)
	}
}

func TestCreateNewCSVDataSourceInvalidPath(t *testing.T) {
	p := ""
	_, err := NewCSVDataSource(p)
	if err == nil {
		t.Errorf("Expected error with invalid path set: %s, as path", p)
	}
}

func TestCreateNewCSVDataStore(t *testing.T) {
	p := "/ondemand-go-bootcamp/poke.csv"
	csvDS, _ := NewCSVDataStore(p)
	if csvDS.csvPath != p {
		t.Errorf("CSVPath was not set, got: %s, want: %s.", csvDS.csvPath, p)
	}
}

func TestCreateNewCSVDataStoreInvalidPath(t *testing.T) {
	p := ""
	_, err := NewCSVDataStore(p)
	if err == nil {
		t.Errorf("Expected error with invalid path set: %s, as path", p)
	}
}
