package opencorporates

import (
	"context"
	"strings"

	"github.com/gocolly/colly"
	"github.com/posolwar/orgcouncil-parse/internal/helpers"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
)

// Получение детальной информации о компаниях в opencorporates
func CompanyDetailConveer(ctx context.Context, in <-chan OpenCorporateCompanyLink) <-chan orgcouncil.CompanyDetailedInfo {
	c := collector.NewCollector()

	out := make(chan orgcouncil.CompanyDetailedInfo)

	go func() {
		var detailedCompany map[string]string

		zipConfirmed := false

		c.OnHTML(".registered_address", func(e *colly.HTMLElement) {
			address := e.Text

			if strings.Contains(address, detailedCompany[helpers.HeaderZip]) {
				zipConfirmed = true
			}
		})

		c.OnHTML(".incorporation_date", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderIncorporationDate] = strings.ToLower(e.Text)
		})

		c.OnHTML(".company_type", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderCompanyType] = strings.ToLower(e.Text)
		})

		c.OnHTML(".jurisdiction", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderJurisdiction] = strings.ToLower(e.Text)
		})

		c.OnHTML(".agent_name", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderAgentName] = strings.ToLower(e.Text)
		})

		c.OnHTML(".agent_address", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderAgentAddress] = strings.ToLower(e.Text)
		})

		c.OnHTML(".trunc8", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderDirectors] = strings.ToLower(e.Text)
		})

		c.OnHTML(".registry_page", func(e *colly.HTMLElement) {
			detailedCompany[helpers.HeaderRegistryPage] = strings.ToLower(e.Text)
		})

		for companyInfo := range in {
			detailedCompany = companyInfo.Info

			if len(companyInfo.URL) == 0 {
				detailedCompany[helpers.HeaderOpencorporatesLink] = helpers.OpenCorporateCompanyLinkError

				out <- detailedCompany
			}

			for _, url := range companyInfo.URL {
				detailedCompany[helpers.HeaderOpencorporatesLink] = url

				c.Visit(url)

				if zipConfirmed {
					out <- detailedCompany
				}
			}
		}

		close(out)
	}()

	return out
}
