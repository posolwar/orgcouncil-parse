package orgcouncil

import (
	"context"

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
			detailedCompany[e.ChildText("th")] = e.ChildText("td")
		})

		for city := range in {
			c.Visit(city.URL)

			helpers.CounterAdd()
			out <- detailedCompany
		}

		close(out)
	}()

	return out
}
