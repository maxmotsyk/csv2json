package integration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	Conv "github.com/maxmotsyk/csv2json/internal/csvconv"
)

func TestCLI(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(dir, "out_data.csv")
	in := filepath.Join(dir, "in_data.json")
	expected := []map[string]string{
		{"name": "Maks", "age": "27", "city": "Kyiv"},
		{"name": "Anna", "age": "25", "city": "Lviv"},
	}

	csvRecord := "name,age,city\nMaks,27,Kyiv\nAnna,25,Lviv\n"
	if err := os.WriteFile(out, []byte(csvRecord), 0o644); err != nil {
		t.Fatal(err.Error())
	}

	csvObj := Conv.CsvData{
		Path: out,
	}

	if err := csvObj.GetRecords(); err != nil {
		t.Fatalf("ReadCSV: %v", err)
	}

	jsonFile := Conv.JsonData{
		Path: in,
	}

	if err := jsonFile.MakeRecords(csvObj.Records); err != nil {
		t.Fatalf("MakeJson: %v", err)
	}

	data, err := os.ReadFile(in)

	if err != nil {
		t.Fatalf("read out.json: %v", err)
	}

	var got []map[string]string
	if err := json.Unmarshal(data, &got); err != nil { // ← обов'язково &got
		t.Fatalf("unmarshal: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("mismatch:\n got=%v\nwant=%v", got, expected)
	}

}
