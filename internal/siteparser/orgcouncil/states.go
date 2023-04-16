package orgcouncil

import (
	"context"

	"github.com/posolwar/orgcouncil-parse/internal/helpers"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"

	"github.com/gocolly/colly"
)

type StateInfo struct {
	Name string
	URL  string
}

func StateConveer(ctx context.Context) <-chan StateInfo {
	c := collector.NewCollector()

	stateOutCh := make(chan StateInfo)

	go func() {
		c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
			text := e.Text

			if text != "" {
				stateOutCh <- StateInfo{Name: text, URL: link}
			}
		})

		c.Visit(helpers.AddressOrgCouncil)

		close(stateOutCh)
	}()

	return stateOutCh
}
