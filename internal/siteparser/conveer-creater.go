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
	if err := csvcreater.CreateDir(dirPath); err != nil {
		logrus.Errorf("ошибка создания каталога '%s', подробности: " + err.Error())
		return
	}

	filteredChannels := make([]<-chan orgcouncil.CompanyDetailedInfo, 0, channelsCount)

	// Это каналы, который работают с небольшим кол-вом информации
	stateCh := orgcouncil.StateConveer(ctx, dirPath, stateFilter)
	stateCh2, fileCh := FileCreateConveer(ctx, dirPath, stateCh)
	cityCh := orgcouncil.CityConveer(ctx, stateCh2)
	companyCh := orgcouncil.CompanyConveer(ctx, cityCh)

	for i := 0; i < channelsCount; i++ {
		// Каналы, который работают с большим кол-во информации
		detailedCh := orgcouncil.CompanyDetailedConveer(ctx, companyCh)
		fileteredCh := orgcouncil.FilteredConveer(ctx, paramFilter, detailedCh)

		filteredChannels = append(filteredChannels, fileteredCh)
	}

	outChan := MergeChannels(filteredChannels...)

	toFileWrite(channelsCount, fileCh, outChan)
}

func MergeChannels(in ...<-chan orgcouncil.CompanyDetailedInfo) <-chan orgcouncil.CompanyDetailedInfo {
	var wg sync.WaitGroup
	var outChannel = make(chan orgcouncil.CompanyDetailedInfo, runtime.NumCPU())

	output := func(channels <-chan orgcouncil.CompanyDetailedInfo) {
		for channel := range channels {
			outChannel <- channel
		}

		wg.Done()
	}

	wg.Add(len(in))
	for _, inChan := range in {
		go output(inChan)
	}

	go func() {
		wg.Wait()
		close(outChannel)
	}()

	return outChannel
}

func toFileWrite(channelsCount int, toWriteFile <-chan csvcreater.CsvToWrite, in <-chan orgcouncil.CompanyDetailedInfo) {
	var wg sync.WaitGroup

	for fileForWriter := range toWriteFile {
		wg.Add(1)
		go toCsvWrite(fileForWriter, &wg, in)
	}

	wg.Wait()
}

func toCsvWrite(fileToWrite csvcreater.CsvToWrite, wg *sync.WaitGroup, ch <-chan orgcouncil.CompanyDetailedInfo) {
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

	fileToWrite.File.Close()
	wg.Done()
}
