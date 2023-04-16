package opencorporates

import (
	"context"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/posolwar/orgcouncil-parse/internal/helpers"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/collector"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
)

type OpenCorporateCompanyLink struct {
	URL  []string
	Info orgcouncil.CompanyDetailedInfo
}

// Получение детальной информации о компаниях в opencorporates
func CompanyListConveer(ctx context.Context, in <-chan orgcouncil.CompanyDetailedInfo) <-chan OpenCorporateCompanyLink {
	c := collector.NewCollector()

	out := make(chan OpenCorporateCompanyLink)

	go func() {
		detailedCompany := OpenCorporateCompanyLink{}

		c.OnHTML(".company_search_result", func(e *colly.HTMLElement) {
			detailedCompany.URL = append(detailedCompany.URL, e.Request.AbsoluteURL(e.ChildAttr(".company_search_result", "href")))
		})

		for city := range in {
			if orgName, ok := city[helpers.HeaderOrganizationName]; ok {
				detailedCompany.Info = city
				generatedLink := opencorporateLinkGenerator(orgName)

				c.Visit(generatedLink)
			}

			out <- detailedCompany
		}

		close(out)
	}()

	return out
}

// Получение ссылки с введенным именем в строке url
func opencorporateLinkGenerator(name string) string {
	// Разбиваем имя на слова
	wordsArray := strings.Fields(name)

	// Соединяем слова с помощью плюса, для верного запроса поиска
	return fmt.Sprintf(helpers.TemplateLinkOpenCorporateSearch, strings.Join(wordsArray, "+"))
}
