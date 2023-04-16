package helpers

const (
	CountOfStates = 56

	// headers
	HeaderOrgcouncilLink              = "orgcouncil link"
	HeaderOrganizationName            = "organization name"
	HeaderTaxPeriodAssetIncomeRevenue = "tax periodassetincomerevenue"
	HeaderZip                         = "zip"
	HeaderIncorporationDate           = "corporation date"
	HeaderCompanyType                 = "company type"
	HeaderJurisdiction                = "jurisdiction"
	HeaderAgentName                   = "agent name"
	HeaderAgentAddress                = "agent address"
	HeaderDirectors                   = "directors"
	HeaderRegistryPage                = "registry page"
	HeaderOpencorporatesLink          = "opencorporates link"

	// Errors
	OpenCorporateCompanyLinkError = "Opencorporate link not found"

	// Allowed domains
	OrgAllowedDomain = "www.orgcouncil.com"

	AddressOrgCouncil    = "https://www.orgcouncil.com/"
	AddressOpenCorporate = "https://opencorporates.com"

	// LinkTemplate
	TemplateLinkOpenCorporateSearch = "https://opencorporates.com/companies?q=%s"
)
