package siteparser

import (
	"context"
	"encoding/csv"
	"sort"
	"sync"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/opencorporates"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
	"github.com/sirupsen/logrus"
)

func CreateConveer(ctx context.Context, csv *csv.Writer, channelsCount int, filter map[string]string) {
	var wg sync.WaitGroup

	// Это каналы, который работают с небольшим кол-вом информации
	stateCh := orgcouncil.StateConveer(ctx)
	cityCh := orgcouncil.CityConveer(ctx, stateCh)

	// Каналы, который работают с большим кол-во информации
	for i := 0; i < channelsCount; i++ {
		companyCh := orgcouncil.CompanyConveer(ctx, cityCh)
		detailedCh := orgcouncil.CompanyDetailedConveer(ctx, companyCh)
		fileteredCh := orgcouncil.FilteredConveer(ctx, filter, detailedCh)
		openCorpListCh := opencorporates.CompanyListConveer(ctx, fileteredCh)
		openCorpDetailCh := opencorporates.CompanyDetailConveer(ctx, openCorpListCh)

		wg.Add(1)
		go toCsvWrite(csv, &wg, openCorpDetailCh)
	}

	wg.Wait()
}

func toCsvWrite(csv *csv.Writer, wg *sync.WaitGroup, ch <-chan orgcouncil.CompanyDetailedInfo) {
	for detailedInfo := range ch {
		csv.Write([]string{"--------------------------------"})
		slice := make([][]string, 0, len(detailedInfo))

		for name, value := range detailedInfo {
			slice = append(slice, []string{name, value})
		}

		sort.Slice(slice, func(i, j int) bool {
			return slice[i][0] < slice[j][0]
		})

		for _, sliceValue := range slice {
			err := csv.Write(sliceValue)
			if err != nil {
				logrus.Errorf("value %v, err: %s", sliceValue, err.Error())
			}
		}
	}

	wg.Done()
}
