package quotesrepo

import (
	"quote/api/app/internal/core/domain"

	"gorm.io/gorm"
)

type QuotesRepo struct {
	db *gorm.DB
}

func NewQuotesRepo(db *gorm.DB) *QuotesRepo {
	return &QuotesRepo{
		db: db,
	}
}

func (q *QuotesRepo) Save(quote domain.Quote) (domain.Quote, error) {
	panic("implement me")
}

func (q *QuotesRepo) GetAllQuotes() ([]domain.Quote, error) {
	panic("implement me")
}

func (q *QuotesRepo) GetQuote(id string) (domain.Quote, error) {
	panic("implement me")
}

func (q *QuotesRepo) UpdateQuote(quote domain.Quote) (domain.Quote, error) {
	panic("implement me")

}
