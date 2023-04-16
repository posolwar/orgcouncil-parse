package siteparser

import (
	"context"
	"encoding/csv"
	"sync"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
	"github.com/sirupsen/logrus"
)

func CreateConveer(ctx context.Context, csv *csv.Writer, channelsCount int) {
	var wg sync.WaitGroup

	stateCh := orgcouncil.StateConveer(ctx)
	cityCh := orgcouncil.CityConveer(ctx, stateCh)

	for i := 0; i < channelsCount; i++ {
		companyCh := orgcouncil.CompanyConveer(ctx, cityCh)
		detailedCh := orgcouncil.CompanyDetailedConveer(ctx, companyCh)
		fileteredCh := orgcouncil.FilteredConveer(ctx, map[string]string{"NTEE Code": "T11"}, detailedCh)

		wg.Add(1)
		go func() {
			for detailedInfo := range fileteredCh {
				logrus.Print(detailedInfo["Organization Name"])
				csv.Write([]string{"--------------------------------"})
				for name, value := range detailedInfo {
					err := csv.Write([]string{name, value})
					if err != nil {
						logrus.Errorf("name: %s, key %s, err: %s", name, value, err.Error())
					}
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
