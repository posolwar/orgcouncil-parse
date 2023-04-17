package orgcouncil

import (
	"context"
)

func Filtration(fieldName, fieldValue string) bool {
	return false
}

// Фильтрация конвеера
func FilteredConveer(ctx context.Context, filterParams map[string]string, in <-chan CompanyDetailedInfo) <-chan CompanyDetailedInfo {
	out := make(chan CompanyDetailedInfo)

	go func() {
		for detailInfo := range in {
			if len(filterParams) == 0 {
				out <- detailInfo
				continue
			}

			if !isFilterConfirmed(filterParams, detailInfo) {
				continue
			}

			out <- detailInfo
		}

		close(out)
	}()

	return out
}

func isFilterConfirmed(filterParams map[string]string, detail CompanyDetailedInfo) bool {
	for paramName, paramValue := range filterParams {
		if detail[paramName] != paramValue {
			return false
		}
	}

	return true
}
