package data

import (
	"encoding/csv"
	"errors"
	"os"
)

type csvDataSource struct {
	csvPath string
}

type CsvDataSource interface {
	ReadCollection() ([][]string, error)
}

func NewCSVDataSource(csvPath string) (csvDataSource, error) {
	if csvPath == "" {
		return csvDataSource{csvPath}, errors.New("No CSV path sent")
	}
	return csvDataSource{csvPath}, nil
}

func (c csvDataSource) ReadCollection() ([][]string, error) {
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

type CsvDataStore interface {
	SaveRecord(record []string) error
	SaveRecords(record [][]string) error
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

func (c csvDataStore) SaveRecords(records [][]string) error {
	file, err := os.OpenFile(c.csvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(records)
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
