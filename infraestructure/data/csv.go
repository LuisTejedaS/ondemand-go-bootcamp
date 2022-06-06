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
	CSVPath() string
}

func NewCSVDataSource(csvPath string) (CsvDataSource, error) {
	if csvPath == "" {
		return csvDataSource{csvPath}, errors.New("No CSV path sent")
	}
	return csvDataSource{csvPath}, nil
}

func (c csvDataSource) CSVPath() string {
	return c.csvPath
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

	if len(records) > 0 {
		records = records[1:]
	}
	return records, nil
}

type CsvDataStore interface {
	SaveRecord(record []string) error
	SaveRecords(record [][]string) error
	DeleteRecords() error
	CSVPath() string
}

type csvDataStore struct {
	csvPath string
}

func NewCSVDataStore(csvPath string) (CsvDataStore, error) {
	if csvPath == "" {
		return csvDataStore{csvPath}, errors.New("No CSV path sent")
	}
	return csvDataStore{csvPath}, nil
}

func (c csvDataStore) CSVPath() string {
	return c.csvPath
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

func (c csvDataStore) DeleteRecords() error {
	if err := os.Truncate(c.csvPath, 0); err != nil {
		return err
	}
	return nil
}
