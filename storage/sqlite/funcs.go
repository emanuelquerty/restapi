package sqlite

import "strings"

func BuildUpdateColumns(updates map[string]any) (names string, values []any) {
	var builder strings.Builder
	for colName, colValue := range updates {
		if colName == "id" || colValue == "" {
			continue
		}
		builder.WriteString(colName)
		builder.WriteString("=?, ")
		values = append(values, colValue)
	}

	names = builder.String()[:builder.Len()-2]
	return
}
