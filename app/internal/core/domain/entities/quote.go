package entities

import (
	"quote/api/app/internal/errdefs"
	"quote/api/app/pkg/identifier"
	"strings"
)

// Quote is a quote entity
type Quote struct {
	identifier.ID
	Author string
	Quote  string
	BaseEntity
}

// NewQuote returns a new quote entity or an error
func NewQuote(author, quote string) (*Quote, error) {
	if len(quote) == 0 {
		return nil, errdefs.ErrInvalidQuote
	}

	id := identifier.New()

	if author == "" {
		author = "Unknown"
	}

	author = strings.Trim(author, " ")

	return &Quote{
		ID:         id,
		Quote:      quote,
		Author:     author,
		BaseEntity: NewBaseEntity(),
	}, nil
}
