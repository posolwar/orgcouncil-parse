package orgcouncil

import (
	"strconv"

	"github.com/gocolly/colly"
)

type CityInfo struct {
	CityName string
	URL      string
	Count    int
}

// Получаем штат, отдаем информацию о городе
func CityConveer(in chan StateInfo, out chan CityInfo, c *colly.Collector) {
	defer close(in)

	c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
		city := CityInfo{}

		rawCount := e.ChildText(".ar")
		city.Count, _ = strconv.Atoi(rawCount)

		city.URL = e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
		city.CityName = e.ChildText("td > a")

		out <- city
	})

	for state := range in {
		c.Visit(state.URL)
	}
}
