package siteparser

import (
	"context"
	"encoding/csv"
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
		for name, value := range detailedInfo {
			err := csv.Write([]string{name, value})
			if err != nil {
				logrus.Errorf("name: %s, key %s, err: %s", name, value, err.Error())
			}
		}
	}

	wg.Done()
}
