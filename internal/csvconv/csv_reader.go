package csvconv

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvData struct {
	Path    string
	Records [][]string
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
