package csvconv

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
)

type CsvData struct {
	Path    string
	Records [][]string
}

func (c *CsvData) GetRecords() error {
	slog.Info("Start CSV read", "input_file_path:", c.Path)

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

	slog.Info("Finish CSV read", "input_file_path:", c.Path)
	return nil
}
