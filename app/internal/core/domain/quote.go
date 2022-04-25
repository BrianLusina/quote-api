package domain

import (
	"quote/api/app/internal/core/domain/entity"
	"quote/api/app/pkg/identifier"
	"strings"
)

type Quote struct {
	identifier.ID[Quote]
	Author string
	Quote  string
	entity.BaseEntity
}

func NewQuote(author, quote string) (*Quote, error) {
	if len(quote) == 0 {
		return nil, ErrInvalidQuote
	}

	id := identifier.New[Quote]()

	if author == "" {
		author = "Unknown"
	}

	author = strings.Trim(author, " ")

	return &Quote{
		ID:         id,
		Quote:      quote,
		Author:     author,
		BaseEntity: entity.NewBaseEntity(),
	}, nil
}

func (q Quote) Prefix() string {
	return "quote"
}
