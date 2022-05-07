package quotesrepo

import (
	"quote/api/app/internal/core/domain/entity"
	"quote/api/app/internal/repositories/models"
	"quote/api/app/pkg/identifier"
	"quote/api/app/tools/logger"
	"time"

	"gorm.io/gorm"
)

type QuotesRepo struct {
	db  *gorm.DB
	log logger.Logger
}

func NewQuotesRepo(db *gorm.DB) *QuotesRepo {
	log := logger.NewLogger("repositories/quotesrepo")

	return &QuotesRepo{
		db:  db,
		log: log,
	}
}

func (q *QuotesRepo) Save(quote entity.Quote) (entity.Quote, error) {
	newQuote := models.Quote{
		Quote:  quote.Quote,
		Author: quote.Author,
		BaseModel: models.BaseModel{
			Identifier: quote.ID.String(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	result := q.db.Create(&newQuote)

	if result.Error != nil {
		q.log.Errorf("Error saving quote: %v", result.Error)
		return entity.Quote{}, result.Error
	}
	return quote, nil
}

func (q *QuotesRepo) GetAllQuotes() ([]entity.Quote, error) {
	var quotes []models.Quote
	result := q.db.Find(&quotes)

	if result.Error != nil {
		q.log.Errorf("Error quering all quotes: %v", result.Error)
		return nil, result.Error
	}

	var allQuotes []entity.Quote
	for _, quote := range quotes {
		allQuotes = append(allQuotes, entity.Quote{
			ID:     identifier.New().FromString(quote.Identifier),
			Quote:  quote.Quote,
			Author: quote.Author,
			BaseEntity: entity.BaseEntity{
				CreatedAt: quote.CreatedAt,
				UpdatedAt: quote.UpdatedAt,
			},
		})
	}
	return allQuotes, nil
}

func (q *QuotesRepo) GetQuote(id string) (entity.Quote, error) {
	var quote models.Quote
	result := q.db.Where(&models.Quote{BaseModel: models.BaseModel{Identifier: id}}).First(&quote)

	if result.Error != nil {
		q.log.Errorf("Error quering quote: %v", result.Error)
		return entity.Quote{}, result.Error
	}

	return entity.Quote{
		ID:     identifier.New().FromString(quote.Identifier),
		Quote:  quote.Quote,
		Author: quote.Author,
		BaseEntity: entity.BaseEntity{
			CreatedAt: quote.CreatedAt,
			UpdatedAt: quote.UpdatedAt,
		},
	}, nil
}

func (q *QuotesRepo) UpdateQuote(quote entity.Quote) (entity.Quote, error) {
	panic("implement me")

}
