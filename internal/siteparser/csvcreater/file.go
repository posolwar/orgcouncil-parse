package csvcreater

import (
	"encoding/csv"
	"os"
)

type CsvToWrite struct {
	File      *os.File
	CsvWriter *csv.Writer
}

func CreateCsv(directoryPath, fileName string) (*csv.Writer, *os.File, error) {
	file, err := os.Create(directoryPath + string(os.PathSeparator) + fileName + ".csv")
	if err != nil {
		return nil, nil, err
	}

	writer := csv.NewWriter(file)

	writer.Comma = ';'

	return writer, file, nil
}

func CreateDir(directoryPath string) error {
	if directoryPath == "" {
		return nil
	}

	return os.MkdirAll(directoryPath, os.ModePerm)
}
