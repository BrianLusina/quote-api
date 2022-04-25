package interactor

import (
	"quote/api/app/internal/core/contracts"
	"quote/api/app/internal/core/domain"
)

type QuoteInteractor struct {
	quoteRepo contracts.QuoteRepository
}

func NewQuoteInteractor(quoteRepo contracts.QuoteRepository) *QuoteInteractor {
	return &QuoteInteractor{quoteRepo: quoteRepo}
}

func (q *QuoteInteractor) CreateQuote(author, quote string) (*domain.Quote, error) {
	newQuote, err := domain.NewQuote(author, quote)
	if err != nil {
		return nil, err
	}

	_, err = q.quoteRepo.Save(*newQuote)
	if err != nil {
		return nil, err
	}
	return newQuote, nil
}

func (q *QuoteInteractor) GetAllQuotes() ([]domain.Quote, error) {
	allQuotes, err := q.quoteRepo.GetAllQuotes()
	if err != nil {
		return nil, err
	}
	return allQuotes, nil
}
