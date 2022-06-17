package entities

import (
	"quote/api/app/pkg/identifier"
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
	if saying, err := newSaying(quote); err != nil {
		return nil, err
	} else {
		quote = saying.String()
	}

	if quoteAuthor, err := newAuthor(author); err != nil {
		return nil, err
	} else {
		author = quoteAuthor.String()
	}

	id := identifier.New()

	return &Quote{
		ID:         id,
		Quote:      quote,
		Author:     author,
		BaseEntity: NewBaseEntity(),
	}, nil
}
