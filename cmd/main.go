package main

import (
	"context"
	"log"
	"runtime"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/csvcreater"
)

func main() {
	ctx := context.Background()

	csv, file, err := csvcreater.CreateCsv("out")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	siteparser.CreateConveer(ctx, csv, runtime.NumCPU()*5)
}

//56
