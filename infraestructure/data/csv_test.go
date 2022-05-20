package data

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createFile(t *testing.T) string {
	file, fe := os.Create("testdata/poke.csv")
	if fe != nil {
		t.Errorf("could not create file for testing: %s", fe)
	}
	dir, pe := filepath.Abs(filepath.Dir(file.Name()))
	if pe != nil {
		t.Errorf("Error getting file path: %s", pe)
	}
	return dir
}

func populateFile(t *testing.T, p string) {
	record := []string{"1", "Charmander"}
	header := []string{"id", "Name"}
	csvDSt, _ := NewCSVDataStore(p)
	csvDSt.SaveRecord(header)
	err := csvDSt.SaveRecord(record)
	if err != nil {
		t.Errorf("There was an error saving record: %s", err)
	}
}

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

func TestReadCollection(t *testing.T) {
	dir := createFile(t)
	p := dir + "/poke.csv"
	populateFile(t, p)
	csvDS, _ := NewCSVDataSource(p)
	if csvDS.csvPath != p {
		t.Errorf("CSVPath was not set, got: %s, want: %s.", csvDS.csvPath, p)
	}
	pok, err := csvDS.ReadCollection()
	if err != nil {
		t.Errorf("Error reading collection %s", err)
	}
	if len(pok) != 1 {
		t.Errorf("Expected to read: %d pokemons but read %d", 1, len(pok))
	}
	os.Remove("testdata/poke.csv")
}

func TestReadCollectionError(t *testing.T) {
	dir := createFile(t)
	p := dir + "/poka.csv"
	csvDS, _ := NewCSVDataSource(p)
	if csvDS.csvPath != p {
		t.Errorf("CSVPath was not set, got: %s, want: %s.", csvDS.csvPath, p)
	}
	_, err := csvDS.ReadCollection()
	if !strings.Contains(err.Error(), "no such file or directory") {
		t.Errorf("We expect to return error reading unexistent file: %s", err.Error())
	}
}

func TestSaveRecord(t *testing.T) {
	dir := createFile(t)
	p := dir + "/poke.csv"
	record := []string{"1", "Charmander"}
	header := []string{"id", "Name"}
	csvDSt, _ := NewCSVDataStore(p)
	csvDSt.SaveRecord(header)
	err := csvDSt.SaveRecord(record)
	if err != nil {
		t.Errorf("There was an error saving record: %s", err)
	}
	csvDS, _ := NewCSVDataSource(p)
	if csvDS.csvPath != p {
		t.Errorf("CSVPath was not set, got: %s, want: %s.", csvDS.csvPath, p)
	}
	pok, err := csvDS.ReadCollection()
	if err != nil {
		t.Errorf("Error reading collection %s", err)
	}
	if len(pok) != 1 {
		t.Errorf("Expected to read: %d pokemons but read %d", 1, len(pok))
	}
	os.Remove("testdata/poke.csv")
}

func TestSaveRecordError(t *testing.T) {
	dir := createFile(t)
	p := dir + "/poka.csv"
	record := []string{"1", "Charmander"}
	header := []string{"id", "Name"}
	csvDSt, _ := NewCSVDataStore(p)
	csvDSt.SaveRecord(header)
	err := csvDSt.SaveRecord(record)
	if !strings.Contains(err.Error(), "no such file or directory") {
		t.Errorf("We expect to return error reading unexistent file: %s", err.Error())
	}
}

func TestSaveRecords(t *testing.T) {
	dir := createFile(t)
	p := dir + "/poke.csv"
	records := [][]string{{"id", "Name"}, {"1", "bulbasaur"}, {"2", "ivysaur"}}

	csvDSt, _ := NewCSVDataStore(p)
	err := csvDSt.SaveRecords(records)
	if err != nil {
		t.Errorf("There was an error saving record: %s", err)
	}
	csvDS, _ := NewCSVDataSource(p)
	if csvDS.csvPath != p {
		t.Errorf("CSVPath was not set, got: %s, want: %s.", csvDS.csvPath, p)
	}
	pok, err := csvDS.ReadCollection()
	if err != nil {
		t.Errorf("Error reading collection %s", err)
	}
	if len(pok) != 2 {
		t.Errorf("Expected to read: %d pokemons but read %d", 1, len(pok))
	}
	os.Remove("testdata/poke.csv")
}
