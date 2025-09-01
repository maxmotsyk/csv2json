package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	// "strings"
)

type CsvData struct {
	Path    string
	Records [][]string
}

type JsonData struct {
	Path string
}

func (c *CsvData) GetRecords() error {
	f, err := os.Open(c.Path)
	if err != nil {
		return err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()

	if err != nil {
		return fmt.Errorf("Failid read from file %w", err)
	}

	c.Records = records

	return nil
}

func (j *JsonData) MakeRecords(csvData [][]string) error {

	if len(csvData) == 0 {
		return fmt.Errorf("CSV data is empty")
	}

	f, err := os.Create(j.Path)

	if err != nil {
		return fmt.Errorf("failed create file %w", err)
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")

	headers := csvData[0]
	results := make([]map[string]string, 0, len(csvData)-1)

	for i := 1; i < len(csvData); i++ {
		obj := makeMap(headers, csvData[i])
		results = append(results, obj)
	}

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("json encode failed: %w", err)
	}
	return nil
}

func makeMap(headers []string, colums []string) map[string]string {
	newMap := make(map[string]string, len(headers))

	for i, head := range headers {
		if i < len(colums) {
			newMap[head] = colums[i]
		} else {
			newMap[head] = ""
		}
	}

	return newMap
}

func main() {
	inputFile := flag.String("in", "input.csv", "CSV file to read from")
	outputFile := flag.String("out", "output.json", "JSON file to write to")
	flag.Parse()

	csvFile := CsvData{
		Path: *inputFile,
	}
	jsonFile := JsonData{
		Path: *outputFile,
	}

	if err := csvFile.GetRecords(); err != nil {
		fmt.Fprintf(os.Stderr, "read csv error: %v\n", err)
		os.Exit(1)
	}
	if err := jsonFile.MakeRecords(csvFile.Records); err != nil {
		fmt.Fprintf(os.Stderr, "write json error: %v\n", err)
		os.Exit(1)
	}
}
