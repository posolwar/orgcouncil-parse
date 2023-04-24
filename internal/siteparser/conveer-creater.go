package siteparser

import (
	"context"
	"runtime"
	"sort"
	"sync"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/csvcreater"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
	"github.com/sirupsen/logrus"
)

func CreateConveer(ctx context.Context, dirPath string, stateFilter map[string]struct{}, paramFilter map[string]string, channelsCount int) {
	var wg sync.WaitGroup
	var outChannel = make(chan orgcouncil.CompanyDetailedInfo, runtime.NumCPU())

	if err := csvcreater.CreateDir(dirPath); err != nil {
		logrus.Errorf("ошибка создания каталога '%s', подробности: " + err.Error())
		return
	}

	// Это каналы, который работают с небольшим кол-вом информации
	stateCh := orgcouncil.StateConveer(ctx, dirPath, stateFilter)
	stateCh2, fileCh := FileCreateConveer(ctx, dirPath, stateCh)
	cityCh := orgcouncil.CityConveer(ctx, stateCh2)

	for i := 0; i < channelsCount; i++ {
		// Каналы, который работают с большим кол-во информации
		companyCh := orgcouncil.CompanyConveer(ctx, cityCh)
		detailedCh := orgcouncil.CompanyDetailedConveer(ctx, companyCh)
		fileteredCh := orgcouncil.FilteredConveer(ctx, paramFilter, detailedCh)
		// openCorpListCh := opencorporates.CompanyListConveer(ctx, fileteredCh)
		// openCorpDetailCh := opencorporates.CompanyDetailConveer(ctx, openCorpListCh)
		go ToOut(fileteredCh, outChannel)
	}

	toCsvWrite2(channelsCount, &wg, fileCh, outChannel)

	wg.Wait()
}

func ToOut(in <-chan orgcouncil.CompanyDetailedInfo, out chan<- orgcouncil.CompanyDetailedInfo) {
	for inChan := range in {
		out <- inChan
	}
}

func toCsvWrite2(channelsCount int, wg *sync.WaitGroup, toWriteFile <-chan csvcreater.CsvToWrite, in <-chan orgcouncil.CompanyDetailedInfo) {
	for fileForWriter := range toWriteFile {
		wg.Add(1)
		go toCsvWrite(fileForWriter, wg, in)
	}
}

func toCsvWrite(fileToWrite csvcreater.CsvToWrite, wg *sync.WaitGroup, ch <-chan orgcouncil.CompanyDetailedInfo) {
	defer fileToWrite.File.Close()
	defer wg.Done()

	for detailedInfo := range ch {
		fileToWrite.CsvWriter.Write([]string{"--------------------------------"})

		slice := make([][]string, 0, len(detailedInfo))

		for name, value := range detailedInfo {
			slice = append(slice, []string{name, value})
		}

		sort.Slice(slice, func(i, j int) bool {
			return slice[i][0] < slice[j][0]
		})

		for _, sliceValue := range slice {
			err := fileToWrite.CsvWriter.Write(sliceValue)
			if err != nil {
				logrus.Errorf("value %v, err: %s", sliceValue, err.Error())
			}
		}
	}
}
