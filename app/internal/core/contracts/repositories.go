package contracts

import "quote/api/app/internal/core/domain"

type QuoteRepository interface {
	Save(domain.Quote) (domain.Quote, error)
	GetAllQuotes() ([]domain.Quote, error)
	GetQuote(id string) (domain.Quote, error)
	UpdateQuote(domain.Quote) (domain.Quote, error)
}
