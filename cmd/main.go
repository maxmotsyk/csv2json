package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	Conv "github.com/maxmotsyk/csv2json/internal/csvconv"
	"github.com/maxmotsyk/csv2json/internal/logger"
)

func main() {
	inputFile := flag.String("in", "input.csv", "CSV file to read from")
	outputFile := flag.String("out", "output.json", "JSON file to write to")
	flag.Parse()

	err := logger.NewLogger()

	if err != nil {
		fmt.Fprintf(os.Stderr, "logger init failed: %v\n", err.Error())
		os.Exit(1)
	}

	csvFile := Conv.CsvData{
		Path: *inputFile,
	}
	jsonFile := Conv.JsonData{
		Path: *outputFile,
	}

	if err := csvFile.GetRecords(); err != nil {
		fmt.Fprintf(os.Stderr, "read csv error: %v\n", err.Error())
		slog.Error("read csv error: %v\n", err)
		os.Exit(1)
	}

	if err := jsonFile.MakeRecords(csvFile.Records); err != nil {
		fmt.Fprintf(os.Stderr, "write json error: %v\n", err.Error())
		slog.Error("write json error: %v\n", err)
		os.Exit(1)
	}
}
