package orgcouncil

import (
	"context"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
)

type CityInfo struct {
	CityName string
	URL      string
	Count    int
}

// Получаем штат, отдаем информацию о городе
func CityConveer(ctx context.Context, in <-chan StateInfo) <-chan CityInfo {
	c := collector.NewCollector()

	cityInfoOut := make(chan CityInfo)

	go func() {
		c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
			city := CityInfo{}

			rawCount := e.ChildText(".ar")
			city.Count, _ = strconv.Atoi(rawCount)

			city.URL = e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
			city.CityName = e.ChildText("td > a")

			if city.CityName != "" {
				cityInfoOut <- city
			}
		})

		for state := range in {
			c.Visit(state.URL)
		}

		close(cityInfoOut)
	}()

	return cityInfoOut
}
