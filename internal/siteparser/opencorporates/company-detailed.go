package opencorporates

// // Получение детальной информации о компаниях в opencorporates
// func CompanyDetailConveer(ctx context.Context, in <-chan OpenCorporateCompanyLink) <-chan orgcouncil.CompanyDetailedInfo {
// 	c := collector.NewCollector()

// 	out := make(chan orgcouncil.CompanyDetailedInfo)

// 	go func() {
// 		detailedCompany := OpenCorporateCompanyLink{}

// 		c.OnHTML(".search-result", func(e *colly.HTMLElement) {
// 			detailedCompany.URL = append(detailedCompany.URL, e.Request.AbsoluteURL(e.ChildAttr(".company_search_result", "href")))
// 		})

// 		for CompanyLink := range in {
// 			if orgName, ok := CompanyLink.Info[helpers.OrganizationName]; ok {
// 				detailedCompany.Info = CompanyLink
// 				c.Visit(opencorporateLinkGenerator(orgName))
// 			}

// 			if detailedCompany.URL != "" {
// 				out <- detailedCompany
// 			}
// 		}

// 		close(out)
// 	}()

// 	return out
// }
