package util

import "strings"

func StringTrimAll(str string) string {
	return strings.TrimSpace(strings.Trim(str, "\n"))
}
