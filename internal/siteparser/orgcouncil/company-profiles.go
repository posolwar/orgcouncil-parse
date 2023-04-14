package orgcouncil

import "github.com/gocolly/colly"

type CompanyProfileInfo struct {
	ID  string
	URL string
}

func CompanyConveer(in chan *CityInfo, out chan *CompanyProfileInfo, c *colly.Collector) {
	defer close(in)

	c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
		profile := CompanyProfileInfo{}
		profile.URL = e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
		profile.ID = e.ChildText(".nowrap")

		out <- &profile
	})

	c.OnHTML(".ac > .pagination > liL", func(e *colly.HTMLElement) {

	})

	for city := range in {
		c.Visit(city.URL)
	}

}
