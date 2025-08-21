package main

import (
	"flag"
	"fmt"
)

func main() {
	inputFile := flag.String("in", "input.csv", "CSV file to read from")
	outputFile := flag.String("out", "output.json", "JSON file to write to")
	flag.Parse()
	fmt.Println(*outputFile, *inputFile)
}
