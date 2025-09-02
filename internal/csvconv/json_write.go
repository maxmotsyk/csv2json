package csvconv

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonData struct {
	Path string
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
