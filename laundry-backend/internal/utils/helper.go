package utils

import (
	"strconv"
	"strings"
)

func QuerySupport(query string) string {
	count := strings.Count(query, "?")
	for i := 0; i < count; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i+1), 1)
	}
	return query
}
