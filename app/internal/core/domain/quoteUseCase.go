package domain

import (
	"math/rand"
	"quote/api/app/internal/core"
	"quote/api/app/internal/core/domain/entities"
	"time"
)

type QuotesUseCase struct {
	quoteRepo core.QuoteRepository
}

func NewQuotesUseCase(quoteRepo core.QuoteRepository) *QuotesUseCase {
	return &QuotesUseCase{quoteRepo: quoteRepo}
}

func (q *QuotesUseCase) CreateQuote(author, quote string) (*entities.Quote, error) {
	newQuote, err := entities.NewQuote(author, quote)
	if err != nil {
		return nil, err
	}

	_, err = q.quoteRepo.Save(*newQuote)
	if err != nil {
		return nil, err
	}
	return newQuote, nil
}

func (q *QuotesUseCase) GetAllQuotes() ([]entities.Quote, error) {
	allQuotes, err := q.quoteRepo.GetAllQuotes()
	if err != nil {
		return nil, err
	}
	return allQuotes, nil
}

func (q *QuotesUseCase) GetQuote(id string) (entities.Quote, error) {
	quote, err := q.quoteRepo.GetQuote(id)
	if err != nil {
		return entities.Quote{}, err
	}
	return quote, nil
}

func (q *QuotesUseCase) GetRandomQuote() (entities.Quote, error) {
	quotes, err := q.quoteRepo.GetAllQuotes()
	if err != nil {
		return entities.Quote{}, err
	}

	// randomize and pick one quote
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	return randomQuote, nil
}
