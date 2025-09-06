package csvconv

import (
	"reflect"
	"testing"
)

func TestMakeMap(t *testing.T) {
	testTable := []struct {
		name     string
		headers  []string
		columns  []string
		expected map[string]string
	}{
		{
			name:    "Full row",
			headers: []string{"name", "age", "city"},
			columns: []string{"Maks", "27", "Kyiv"},
			expected: map[string]string{
				"name": "Maks",
				"age":  "27",
				"city": "Kyiv",
			},
		},
		{
			name:    "Short row (missing values)",
			headers: []string{"name", "age", "city"},
			columns: []string{"Anna"},
			expected: map[string]string{
				"name": "Anna",
				"age":  "",
				"city": "",
			},
		},
		{
			name:    "Long row (extra values ignored)",
			headers: []string{"name", "age"},
			columns: []string{"Bob", "30", "Lviv"},
			expected: map[string]string{
				"name": "Bob",
				"age":  "30",
			},
		},
	}

	for _, item := range testTable {
		t.Run(item.name, func(t *testing.T) {
			got := MakeMap(item.headers, item.columns)

			if !reflect.DeepEqual(got, item.expected) {
				t.Errorf("got=%v, want=%v", got, item.expected)
			}
		})
	}

}
