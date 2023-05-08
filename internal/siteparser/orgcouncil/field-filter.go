package orgcouncil

import (
	"context"
	"strings"
)

func Filtration(fieldName, fieldValue string) bool {
	return false
}

// Фильтрация конвеера
func FilteredConveer(ctx context.Context, filterParams map[string]string, in <-chan CompanyDetailedInfo) <-chan CompanyDetailedInfo {
	out := make(chan CompanyDetailedInfo)

	go func() {
		var isFiltred bool

		if len(filterParams) > 0 {
			isFiltred = true
		}

		for detailInfo := range in {
			if !isFiltred {
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

func isWildcard(filter string) bool {
	return strings.Contains(filter, "*")
}

func searchWildcard(field, filterWithoutWildcard string) bool {
	return strings.Contains(field, filterWithoutWildcard)
}

func isFilterConfirmed(filterParams map[string]string, detail CompanyDetailedInfo) bool {
	for paramName, paramValue := range filterParams {
		paramWithoutStar := strings.ReplaceAll(paramValue, "*", "")

		if isWildcard(paramValue) {
			if !searchWildcard(detail[paramName], paramWithoutStar) {
				return false
			}
		} else {
			if detail[paramName] != paramValue {
				return false
			}
		}
	}

	return true
}
