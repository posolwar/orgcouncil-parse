package orgcouncil

import (
	"context"

	"github.com/posolwar/orgcouncil-parse/internal/helpers"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
	"github.com/sirupsen/logrus"

	"github.com/gocolly/colly"
)

type StateInfo struct {
	Name string
	URL  string
}

// Конвеер для получения информации по штатам
func StateConveer(ctx context.Context, filtredStates map[string]struct{}) <-chan StateInfo {
	c := collector.NewCollector()

	stateOutCh := make(chan StateInfo)

	go func() {
		var isFiltered bool

		// Записываем что фильтрация нужна или нет
		if len(filtredStates) > 0 {
			isFiltered = true
		}

		c.OnHTML(".table-condensed2 > tbody > tr", func(e *colly.HTMLElement) {
			stateLink := e.Request.AbsoluteURL(e.ChildAttr("td > a", "href"))
			stateName := e.Text

			if isFiltered {
				if _, stateFound := filtredStates[stateName]; stateFound {
					stateOutCh <- StateInfo{Name: stateName, URL: stateLink}
				}
			} else {
				if stateName != "" {
					stateOutCh <- StateInfo{Name: stateName, URL: stateLink}
				}
			}
		})

		err := c.Visit(helpers.AddressOrgCouncil)
		if err != nil {
			logrus.Error(err)
			ctx.Done()
		}

		close(stateOutCh)
	}()

	return stateOutCh
}
