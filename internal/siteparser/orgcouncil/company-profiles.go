package orgcouncil

import (
	"context"

	"github.com/gocolly/colly"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
)

type CompanyProfileInfo struct {
	ID  string
	URL string
}

func CompanyConveer(ctx context.Context, in <-chan CityInfo) <-chan CompanyProfileInfo {
	c := collector.NewCollector()

	out := make(chan CompanyProfileInfo)

	go func() {
		c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
			profile := CompanyProfileInfo{}
			profile.URL = e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
			profile.ID = e.ChildText(".nowrap")

			if profile.ID != "" {
				out <- profile
			}
		})

		c.OnHTML(".ac > .pagination > li:last-child", func(e *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
			c.Visit(link)
		})

		for city := range in {
			c.Visit(city.URL)
		}

		close(out)
	}()

	return out
}
