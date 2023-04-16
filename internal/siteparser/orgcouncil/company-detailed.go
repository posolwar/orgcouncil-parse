package orgcouncil

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/gocolly/colly"

	"github.com/posolwar/orgcouncil-parse/internal/helpers"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
)

type CompanyDetailedInfo map[string]string

// Получение детальной информации о компаниях
func CompanyDetailedConveer(ctx context.Context, in <-chan CompanyProfileInfo) <-chan CompanyDetailedInfo {
	c := collector.NewCollector()

	out := make(chan CompanyDetailedInfo)

	go func() {
		detailedCompany := make(CompanyDetailedInfo)

		c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
			header := strings.ToLower(e.ChildText("th"))
			value := strings.ToLower(e.ChildText("td"))

			if header != "" {
				detailedCompany[header] = value
			} else {
				lenValue := utf8.RuneCountInString(value)

				// Проверяю что это строка с содержанием zip
				if isZipLine(value, lenValue) {
					detailedCompany[helpers.HeaderZip] = value[lenValue-5:]
				}

				if isMonth(value, lenValue) {
					detailedCompany[helpers.HeaderTaxPeriodAssetIncomeRevenue] = detailedCompany[helpers.HeaderTaxPeriodAssetIncomeRevenue] + "\n" + value
				}
			}
		})

		for city := range in {
			c.Visit(city.URL)
			detailedCompany[helpers.HeaderOrgcouncilLink] = city.URL

			helpers.CounterAdd()
			out <- detailedCompany
		}

		close(out)
	}()

	return out
}

func isZipLine(value string, lenValue int) bool {
	return lenValue > 3 && value[0:3] == "All"
}

func isMonth(value string, lenValue int) bool {
	if lenValue > 8 {
		return strings.HasPrefix(value, "January") ||
			strings.HasPrefix(value, "February") ||
			strings.HasPrefix(value, "March") ||
			strings.HasPrefix(value, "April") ||
			strings.HasPrefix(value, "May") ||
			strings.HasPrefix(value, "June") ||
			strings.HasPrefix(value, "July") ||
			strings.HasPrefix(value, "August") ||
			strings.HasPrefix(value, "September") ||
			strings.HasPrefix(value, "October") ||
			strings.HasPrefix(value, "November") ||
			strings.HasPrefix(value, "December")
	}

	return false
}
