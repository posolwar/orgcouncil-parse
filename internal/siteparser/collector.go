package siteparser

import (
	"github.com/gocolly/colly"
)

func NewCollector(allowedDomains ...string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomains...),
	)

	return c
}
