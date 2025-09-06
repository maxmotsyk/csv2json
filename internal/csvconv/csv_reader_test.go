package csvconv

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetRecord(t *testing.T) {
	testTable := []struct {
		name     string
		testData string
		expected [][]string
	}{
		{
			name:     "Simple CSV",
			testData: "name,age,city\nMaks,27,Kyiv\nAnna,25,Lviv\n",
			expected: [][]string{
				{"name", "age", "city"},
				{"Maks", "27", "Kyiv"},
				{"Anna", "25", "Lviv"},
			},
		},
		{
			name:     "Missing column",
			testData: "name,age,city\nBob,30\n",
			expected: [][]string{
				{"name", "age", "city"},
				{"Bob", "30", ""},
			},
		},
		{
			name:     "Extra column",
			testData: "name,age\nOleh,40,Lviv\n",
			expected: [][]string{
				{"name", "age", ""},
				{"Oleh", "40", "Lviv"},
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			csvPath := filepath.Join(dir, "imput.csv")

			if err := os.WriteFile(csvPath, []byte(tc.testData), 0644); err != nil {
				t.Fatal(err)
			}

			csvData := CsvData{
				Path: csvPath,
			}

			err := csvData.GetRecords()

			if err != nil {
				t.Fatalf("ReadCSV error: %v", err)
			}

			if !reflect.DeepEqual(tc.expected, csvData.Records) {
				t.Errorf("mismatch:\n got=%v\nwant=%v", csvData.Records, tc.expected)
			}

		})
	}
}
