package data

import (
	"encoding/csv"
	"os"
)

type CsvDataSource struct {
	csvPath string
}

func NewCSVDataSource(csvPath string) CsvDataSource {
	return CsvDataSource{csvPath}
}

func (c CsvDataSource) ReadCollection() ([][]string, error) {
	file, err := os.Open(c.csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	records = records[1:]
	return records, nil
}

type csvDataStore struct {
	store string
}

func NewCSVDataStore(csvPath string) csvDataStore {
	return csvDataStore{store: csvPath}
}

func (c csvDataStore) SaveRecord(record []string) error {
	file, err := os.OpenFile(c.store, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	err = csvWriter.Write(record)
	if err != nil {
		return err
	}

	csvWriter.Flush()

	err = csvWriter.Error()
	if err != nil {
		return err
	}

	return nil
}
