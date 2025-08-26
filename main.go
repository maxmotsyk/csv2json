package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
		panic(err)
	}

	c.Records = records

	return nil
}

func (j *JsonData) MakeRecords(csvData [][]string) error {
	f, err := os.Create(j.Path)

	if err != nil {
		return err
	}

	defer f.Close()

	encoder := json.NewEncoder(f)

	csvMap := make(map[string]string, len(csvData))

	for _, col := range csvData[0] {
		csvMap[col] = ""

		for _, row := range csvData[1:] {
			for _, item := range row {
				csvMap[col] = item
			}
		}
	}

	encoder.Encode(csvMap)

	return nil
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

	csvFile.GetRecords()
	jsonFile.MakeRecords(csvFile.Records)

	fmt.Println(csvFile.Records)

}
