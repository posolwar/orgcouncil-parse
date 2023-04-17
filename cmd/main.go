package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"runtime"
	"strings"
	"unicode/utf8"

	"github.com/posolwar/orgcouncil-parse/internal/siteparser"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/csvcreater"
	"github.com/posolwar/orgcouncil-parse/internal/siteparser/filters"
	"github.com/sirupsen/logrus"
)

var (
	CountOfChannels int
	CsvComma        string
	States          string
	FilterFilePath  string
	OutFileName     string
)

func init() {
	flag.IntVar(&CountOfChannels, "channels", 5, "Кол-во используемых циклов, число также умножается на кол-во ядер. Чем выше, тем больше нагрузка на проц.")

	flag.StringVar(&CsvComma, "csv-comma", ":", "Разделитель, используемый для csv.")
	flag.StringVar(&OutFileName, "out-file-name", "out", "Имя файла, который будет создан для перечисления ответов.")
	flag.StringVar(&FilterFilePath, "file", "", "Путь к файлу, содержащего параметры в формате json для фильтрации по ним.")

	flag.StringVar(&States, "state", "", "Штат, используемый для поиска. Если не указан, то поиск выполняется по всем штатам. Можно указать несколько штатов через запятую.")
}

func main() {
	flag.Parse()

	if err := flagsValid(); err != nil {
		logrus.Fatal(err.Error())
	}

	log.Println("Кол-во запущенных циклов: ", runtime.NumCPU()*CountOfChannels)

	ctx := context.Background()

	csv, file, err := csvcreater.CreateCsv("out")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	paramFilter, err := filters.GetFiltersFromFile(FilterFilePath)
	if err != nil {
		log.Fatal(err)
	}

	stateFilter := getStateFilter(States)

	siteparser.CreateConveer(ctx, stateFilter, paramFilter, csv, runtime.NumCPU()*CountOfChannels)
}

// Получаем список штатов для ограничения поиска
func getStateFilter(rawStates string) map[string]struct{} {
	states := strings.Split(rawStates, ",")

	mapStates := make(map[string]struct{}, len(states))

	for _, state := range states {
		if state != "" {
			mapStates[state] = struct{}{}
		}
	}

	return mapStates
}

func flagsValid() error {
	commaLen := utf8.RuneCountInString(CsvComma)

	// channels
	if CountOfChannels == 0 {
		return errors.New("количество каналов должно быть более единицы")
	}

	// comma
	if commaLen > 1 || commaLen == 0 {
		return errors.New("разделитель должен состоять из одного символа")
	}

	// out file
	if utf8.RuneCountInString(OutFileName) == 0 {
		return errors.New("вы не указали имя файла, в который будет выводиться ответ")
	}

	return nil
}
