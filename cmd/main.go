package main

import (
	"flag"
	"fmt"
	Conv "github.com/maxmotsyk/csv2json/internal/csvconv"
	"log/slog"
	"os"
)

func main() {
	inputFile := flag.String("in", "input.csv", "CSV file to read from")
	outputFile := flag.String("out", "output.json", "JSON file to write to")
	flag.Parse()

	f, _ := os.Create("slog.json")
	logger := slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// міняємо тільки атрибут "time"
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				// формат у стилі "2006-01-02 15:04:05"
				return slog.Attr{
					Key:   slog.TimeKey,
					Value: slog.StringValue(t.Format("2006-01-02 15:04:05")),
				}
			}
			return a
		},
	}))
	slog.SetDefault(logger)

	csvFile := Conv.CsvData{
		Path: *inputFile,
	}
	jsonFile := Conv.JsonData{
		Path: *outputFile,
	}

	if err := csvFile.GetRecords(); err != nil {
		fmt.Fprintf(os.Stderr, "read csv error: %v\n", err)
		logger.Error("read csv error: %v\n", err)
		os.Exit(1)
	}
	if err := jsonFile.MakeRecords(csvFile.Records); err != nil {
		fmt.Fprintf(os.Stderr, "write json error: %v\n", err)
		logger.Error("write json error: %v\n", err)
		os.Exit(1)
	}
}
