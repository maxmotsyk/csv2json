package csvconv

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestMakeRecords(t *testing.T) {
	testTable := []struct {
		name     string
		testData [][]string
		expected []map[string]string
	}{
		{
			name: "Full data",
			testData: [][]string{
				{"name", "age", "city"},
				{"Maks", "27", "Kyiv"},
				{"Anna", "25", "Lviv"},
			},
			expected: []map[string]string{
				{"name": "Maks", "age": "27", "city": "Kyiv"},
				{"name": "Anna", "age": "25", "city": "Lviv"},
			},
		},
		{
			name: "Short row (missing values)",
			testData: [][]string{
				{"name", "age", "city"},
				{"Bob"},
			},
			expected: []map[string]string{
				{"name": "Bob", "age": "", "city": ""},
			},
		},
		{
			name: "Long row (extra values ignored)",
			testData: [][]string{
				{"name", "age"},
				{"Oleh", "30", "extra"},
			},
			expected: []map[string]string{
				{"name": "Oleh", "age": "30"},
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			// 1) окрема тимчасова директорія для кожного під-тесту
			dir := t.TempDir()
			outPath := filepath.Join(dir, "out.json")

			// 2) викликаємо прод-функцію
			j := JsonData{Path: outPath}
			if err := j.MakeRecords(tc.testData); err != nil {
				t.Fatalf("MakeRecords error: %v", err)
			}

			// 3) читаємо файл, який щойно записали
			data, err := os.ReadFile(outPath)
			if err != nil {
				t.Fatalf("read out.json: %v", err)
			}

			// 4) розпарсимо JSON у структуру і порівняємо
			var got []map[string]string
			if err := json.Unmarshal(data, &got); err != nil { // ← обов'язково &got
				t.Fatalf("unmarshal: %v", err)
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("mismatch:\n got=%v\nwant=%v", got, tc.expected)
			}
		})
	}
}
