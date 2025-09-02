package csvconv

import (
	"log/slog"
)

func makeMap(headers []string, colums []string) map[string]string {
	newMap := make(map[string]string, len(headers))

	for i, head := range headers {
		if i < len(colums) {
			newMap[head] = colums[i]
		} else {
			newMap[head] = ""
		}
	}

	slog.Info("Conversion of CSV to MAP has ended")
	return newMap
}
