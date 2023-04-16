package orgcouncil

import (
	"context"

	"github.com/posolwar/orgcouncil-parse/internal/helpers"
)

func Filtration(fieldName, fieldValue string) bool {
	return false
}

func FilteredConveer(ctx context.Context, filterParams map[string]string, in <-chan CompanyDetailedInfo) <-chan CompanyDetailedInfo {
	out := make(chan CompanyDetailedInfo)

	i := 0

	go helpers.Counter("filter", &i)

	go func() {
		for detailInfo := range in {
			var skipThisInfo bool

			i++

			for paramName, paramValue := range filterParams {
				if detailInfo[paramName] != paramValue {
					skipThisInfo = true
					break
				}
			}

			if skipThisInfo {
				continue
			}

			out <- detailInfo
		}

		close(out)
	}()

	return out
}
