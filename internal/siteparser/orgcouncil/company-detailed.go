package orgcouncil

import (
	"context"

	"github.com/gocolly/colly"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
)

type CompanyDetailedInfo map[string]string

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

			CounterAdd()
			out <- detailedCompany
		}

		close(out)
	}()

	return out
}
