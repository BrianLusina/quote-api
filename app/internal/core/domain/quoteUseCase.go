package domain

import (
	"math/rand"
	"quote/api/app/internal/core/contracts"
	"quote/api/app/internal/core/domain/entity"
	"time"
)

type QuotesUseCase struct {
	quoteRepo contracts.QuoteRepository
}

func NewQuotesUseCase(quoteRepo contracts.QuoteRepository) *QuotesUseCase {
	return &QuotesUseCase{quoteRepo: quoteRepo}
}

func (q *QuotesUseCase) CreateQuote(author, quote string) (*entity.Quote, error) {
	newQuote, err := entity.NewQuote(author, quote)
	if err != nil {
		return nil, err
	}

	_, err = q.quoteRepo.Save(*newQuote)
	if err != nil {
		return nil, err
	}
	return newQuote, nil
}

func (q *QuotesUseCase) GetAllQuotes() ([]entity.Quote, error) {
	allQuotes, err := q.quoteRepo.GetAllQuotes()
	if err != nil {
		return nil, err
	}
	return allQuotes, nil
}

func (q *QuotesUseCase) GetQuote(id string) (entity.Quote, error) {
	quote, err := q.quoteRepo.GetQuote(id)
	if err != nil {
		return entity.Quote{}, err
	}
	return quote, nil
}

func (q *QuotesUseCase) GetRandomQuote() (entity.Quote, error) {
	quotes, err := q.quoteRepo.GetAllQuotes()
	if err != nil {
		return entity.Quote{}, err
	}

	// randomize and pick one quote
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	return randomQuote, nil
}
