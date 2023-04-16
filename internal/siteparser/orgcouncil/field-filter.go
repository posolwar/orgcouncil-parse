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
			var skipThisDetailed bool

			for paramName, paramValue := range filterParams {
				if detailInfo[paramName] != paramValue {
					skipThisDetailed = true
					break
				}
			}

			if skipThisDetailed {
				continue
			}

			out <- detailInfo
		}

		close(out)
	}()

	return out
}
