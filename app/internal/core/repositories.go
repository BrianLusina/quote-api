package core

import (
	"quote/api/app/internal/core/entities"
)

type QuoteRepository interface {
	Save(entities.Quote) (entities.Quote, error)
	GetAllQuotes() ([]entities.Quote, error)
	GetQuote(id string) (entities.Quote, error)
	UpdateQuote(entities.Quote) (entities.Quote, error)
}
