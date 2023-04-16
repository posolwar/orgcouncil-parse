package collector

import (
	"github.com/gocolly/colly"
	"github.com/posolwar/orgcouncil-parse/internal/helpers"
)

func NewCollector() *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(helpers.OrgAllowedDomain),
		// colly.Async(true),
	)
}
