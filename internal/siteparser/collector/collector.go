package collector

import (
	"github.com/gocolly/colly"
	"github.com/posolwar/orgcouncil-parse/internal/helpers"
)

// Клиент, который будет ходить по сайту
func NewCollector() *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(helpers.OrgAllowedDomain),
	)
}

func NewSyncCollector() *colly.Collector {
	return colly.NewCollector(
		colly.CacheDir("./cache-dir"),
		colly.AllowedDomains(helpers.OrgAllowedDomain),
		colly.Async(true),
	)
}
