package tools

import "strings"

// StandardizeSpaces returns a string with all spaces replaced by a single space
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
