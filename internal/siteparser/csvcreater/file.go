package csvcreater

import (
	"encoding/csv"
	"os"
)

func CreateCsv(fileName string, csvComma rune) (*csv.Writer, *os.File, error) {
	file, err := os.Create(fileName + ".csv")
	if err != nil {
		return nil, nil, err
	}

	writer := csv.NewWriter(file)

	writer.Comma = csvComma

	return writer, file, nil
}
