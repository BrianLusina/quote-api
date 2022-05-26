package entities

import (
	"fmt"
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
		quoteValue := fmt.Sprintf("%s", saying)
		quote = quoteValue
	}

	if quoteAuthor, err := newAuthor(author); err != nil {
		return nil, err
	} else {
		authorValue := fmt.Sprintf("%s", quoteAuthor)
		author = authorValue
	}

	id := identifier.New()

	return &Quote{
		ID:         id,
		Quote:      quote,
		Author:     author,
		BaseEntity: NewBaseEntity(),
	}, nil
}
