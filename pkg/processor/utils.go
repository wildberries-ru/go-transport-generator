package processor

import "strings"

func keys2String(v map[string]string) string {
	keys := make([]string, 0, len(v))
	for key := range v {
		keys = append(keys, key)
	}
	return strings.Join(keys, ", ")
}
