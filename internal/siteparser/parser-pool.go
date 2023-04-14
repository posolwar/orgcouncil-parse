package siteparser

import (
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/opencorporates"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/orgcouncil"
)

type ParserPool struct {
	StateCh                 chan orgcouncil.StateInfo
	CityCh                  chan orgcouncil.CityInfo
	OrgCouncilCompanyCh     chan orgcouncil.CompanyInfo
	OpenCorporatesCompanyCh chan opencorporates.CompanyInfo
}

func NewParsePool() *ParserPool {
	return &ParserPool{
		StateCh:                 make(chan orgcouncil.StateInfo),
		CityCh:                  make(chan orgcouncil.CityInfo),
		OrgCouncilCompanyCh:     make(chan orgcouncil.CompanyInfo),
		OpenCorporatesCompanyCh: make(chan opencorporates.CompanyInfo),
	}
}
