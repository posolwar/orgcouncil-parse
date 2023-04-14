package orgcouncil

import "github.com/gocolly/colly"

type CompanyDetailedInfo map[string]string

func CompanyDetailedConveer(in chan *CompanyProfileInfo, c *colly.Collector) <-chan CompanyDetailedInfo {
	defer close(in)

	c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
		profile := CompanyProfileInfo{}
		profile.URL = e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
		profile.ID = e.ChildText(".nowrap")

		out <- &profile
	})

	for city := range in {
		c.Visit(city.URL)
	}
}
