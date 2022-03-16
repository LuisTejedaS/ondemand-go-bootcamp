package data

import (
	"encoding/csv"
	"errors"
	"os"
)

type CsvDataSource struct {
	csvPath string
}

func NewCSVDataSource(csvPath string) (CsvDataSource, error) {
	if csvPath == "" {
		return CsvDataSource{csvPath}, errors.New("No CSV path sent")
	}
	return CsvDataSource{csvPath}, nil
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
	csvPath string
}

func NewCSVDataStore(csvPath string) (csvDataStore, error) {
	if csvPath == "" {
		return csvDataStore{csvPath}, errors.New("No CSV path sent")
	}
	return csvDataStore{csvPath}, nil
}

func (c csvDataStore) SaveRecord(record []string) error {
	file, err := os.OpenFile(c.csvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
