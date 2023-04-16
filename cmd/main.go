package main

import (
	"context"
	"flag"
	"log"
	"runtime"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/csvcreater"
)

var CountOfChannels int

func init() {
	flag.IntVar(&CountOfChannels, "channels", 1, "кол-во используемых каналов")
}

func main() {
	flag.Parse()

	log.Println("Кол-во запущенных горутин: ", runtime.NumCPU()*CountOfChannels)

	ctx := context.Background()

	csv, file, err := csvcreater.CreateCsv("out")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	siteparser.CreateConveer(ctx, csv, runtime.NumCPU()*CountOfChannels, map[string]string{"ntee code": "t11"})
}

//56
