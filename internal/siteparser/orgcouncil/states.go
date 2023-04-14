package orgcouncil

import (
	"github.com/posolwar/orgcouncil-parse/internal/helpers"

	"github.com/gocolly/colly"
)

type StateInfo struct {
	Name string
	URL  string
}

func StateConveer(out chan<- StateInfo, c *colly.Collector) {
	c.OnHTML(".table-condensed2 > tbody > tr > td", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		text := e.Text

		if text != "" {
			out <- StateInfo{Name: text, URL: link}
		}
	})

	c.Visit(helpers.AddressOrgCouncil)
}
