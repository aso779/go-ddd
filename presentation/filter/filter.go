package filter

import "strings"

func FieldName(parents []string, field string) string {
	return strings.Join(append(parents, field), ".")
}
