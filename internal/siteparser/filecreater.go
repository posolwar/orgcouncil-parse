package siteparser

import (
	"context"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/csvcreater"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
	"github.com/sirupsen/logrus"
)

func FileCreateConveer(ctx context.Context, dirPath string, states <-chan orgcouncil.StateInfo) (<-chan orgcouncil.StateInfo, <-chan csvcreater.CsvToWrite) {
	outStates := make(chan orgcouncil.StateInfo)
	outFile := make(chan csvcreater.CsvToWrite)

	go func() {
		for state := range states {
			csvWriter, file, err := csvcreater.CreateCsv(dirPath, state.Name)
			if err != nil {
				logrus.Errorf("ошибка создания csv файла для вливания информации о штате %s. Подробности: %s", state.Name, err)
			}

			outFile <- csvcreater.CsvToWrite{File: file, CsvWriter: csvWriter}
			outStates <- state
		}

		close(outStates)
		close(outFile)
	}()

	return outStates, outFile
}
