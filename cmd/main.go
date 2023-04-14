package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/csvcreater"
)

func main() {
	collector := siteparser.NewCollector("www.orgcouncil.com")

	csv, file, err := csvcreater.CreateCsv("out")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	parserPool := siteparser.NewParsePool()

	// states, err := orgcouncil.GetStates(collector, "https://www.orgcouncil.com/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cities, err := orgcouncil.GetStatesOrganizations(states, collector)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(cities)
}

//56
