package siteparser

import (
	"context"
	"sort"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
)

func SortConveer(ctx context.Context, in <-chan orgcouncil.CompanyDetailedInfo) <-chan [][]string {
	out := make(chan [][]string)

	go func() {
		for detailInfo := range in {
			sortedDetails := make([][]string, 0, len(detailInfo))

			for key, value := range detailInfo {
				sortedDetails = append(sortedDetails, []string{key, value})
			}

			sort.SliceStable(sortedDetails, func(i, j int) bool {
				return sortedDetails[i][0] < sortedDetails[j][0]
			})

			out <- sortedDetails
		}

		close(out)
	}()

	return out
}
