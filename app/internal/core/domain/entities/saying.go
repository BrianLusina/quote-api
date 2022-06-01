package entities

import (
	"quote/api/app/internal/errdefs"
	"quote/api/app/tools"
	"regexp"
	"strings"
)

type saying struct {
	value string
}

const (
	sayingRegexPattern = `[a-zA-Z0-9]+`
)

var sayingRegex = regexp.MustCompile(sayingRegexPattern)

// newSaying returns a new saying entity or an error
func newSaying(value string) (*saying, error) {
	matches := sayingRegex.FindAllString(value, -1)

	if len(matches) == 0 {
		return nil, errdefs.ErrInvalidQuote
	} else {
		allMatches := strings.Join(matches, " ")
		value = tools.StandardizeSpaces(allMatches)

		// capitalize the first letter of the string only
		firstCharacter := strings.ToUpper(value[:1])
		value = firstCharacter + value[1:]
	}

	return &saying{value: value}, nil
}

// String returns the author type as a string
func (s *saying) String() string {
	return s.value
}
