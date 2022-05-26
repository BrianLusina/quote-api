package entities

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	authorRegexPattern = `[a-zA-Z\s]+`
)

var authorRegex = regexp.MustCompile(authorRegexPattern)

type author struct {
	value string
}

// newAuthor returns a new author entity or an error
func newAuthor(value string) (*author, error) {
	matches := authorRegex.FindAllString(value, -1)

	if len(matches) == 0 {
		value = "Unknown"
	} else {
		value = strings.Title(strings.Join(matches, " "))
	}

	return &author{value: value}, nil
}

// String returns the author type as a string
func (a *author) String() string {
	return fmt.Sprintf("%s", a.value)
}
