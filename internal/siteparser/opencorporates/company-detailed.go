package opencorporates

import (
	"context"
	"strings"

	"github.com/gocolly/colly"
	"github.com/posolwar/orgcouncil-parse/internal/helpers"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
)

type Details struct {
	OrgcouncilMap     map[string]string
	OpencorporateData [][]string
}

// Получение детальной информации о компаниях в opencorporates
func CompanyDetailConveer(ctx context.Context, in <-chan OpenCorporateCompanyLink) <-chan Details {
	c := collector.NewCollector()

	out := make(chan Details)

	go func() {
		detailedCompany := Details{}

		zipConfirmed := false

		c.OnHTML(".registered_address", func(e *colly.HTMLElement) {
			address := e.Text

			if strings.Contains(address, detailedCompany.OrgcouncilMap[helpers.HeaderZip][0:2]) {
				zipConfirmed = true
			}
		})

		c.OnHTML(".incorporation_date", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderIncorporationDate, strings.ToLower(e.Text)})
		})

		c.OnHTML(".company_type", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderCompanyType, strings.ToLower(e.Text)})
		})

		c.OnHTML(".jurisdiction", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderJurisdiction, strings.ToLower(e.Text)})
		})

		c.OnHTML(".agent_name", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderAgentName, strings.ToLower(e.Text)})
		})

		c.OnHTML(".agent_address", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderAgentAddress, strings.ToLower(e.Text)})
		})

		c.OnHTML(".trunc8", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderDirectors, strings.ToLower(e.Text)})
		})

		c.OnHTML(".registry_page", func(e *colly.HTMLElement) {
			detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderRegistryPage, strings.ToLower(e.Text)})
		})

		for companyInfo := range in {
			detailedCompany.OpencorporateData = make([][]string, 0, 8)
			detailedCompany.OrgcouncilMap = companyInfo.Info

			if len(companyInfo.URL) == 0 {
				detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderOpencorporatesLink, helpers.OpenCorporateCompanyLinkError})

				out <- detailedCompany
			}

			for _, url := range companyInfo.URL {
				detailedCompany.OpencorporateData = append(detailedCompany.OpencorporateData, []string{helpers.HeaderOpencorporatesLink, url})

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
