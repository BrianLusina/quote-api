package quotesrepo

import (
	"quote/api/app/internal/core/entities"
	"quote/api/app/internal/repositories/models"
	"quote/api/app/pkg/identifier"
	"quote/api/app/tools/logger"

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

func (q *QuotesRepo) Save(quote entities.Quote) (entities.Quote, error) {
	newQuote := models.Quote{
		Quote:  quote.Quote,
		Author: quote.Author,
		BaseModel: models.BaseModel{
			Identifier: quote.ID.String(),
		},
	}

	result := q.db.Create(&newQuote)

	if result.Error != nil {
		q.log.Errorf("Error saving quote: %v", result.Error)
		return entities.Quote{}, result.Error
	}
	return quote, nil
}

func (q *QuotesRepo) GetAllQuotes() ([]entities.Quote, error) {
	var quotes []models.Quote

	if err := q.db.Find(&quotes).Error; err != nil {
		q.log.Errorf("Error quering all quotes: %v", err)
		return nil, err
	}

	allQuotes := []entities.Quote{}
	for _, quote := range quotes {
		allQuotes = append(allQuotes, entities.Quote{
			ID:     identifier.New().FromString(quote.Identifier),
			Quote:  quote.Quote,
			Author: quote.Author,
			BaseEntity: entities.BaseEntity{
				CreatedAt: quote.BaseModel.CreatedAt,
				UpdatedAt: quote.BaseModel.UpdatedAt,
			},
		})
	}

	return allQuotes, nil
}

func (q *QuotesRepo) GetQuote(id string) (entities.Quote, error) {
	var quote models.Quote
	result := q.db.Where(&models.Quote{BaseModel: models.BaseModel{Identifier: id}}).First(&quote)

	if result.Error != nil {
		q.log.Errorf("Error quering quote: %v", result.Error)
		return entities.Quote{}, result.Error
	}

	return entities.Quote{
		ID:     identifier.New().FromString(quote.Identifier),
		Quote:  quote.Quote,
		Author: quote.Author,
		BaseEntity: entities.BaseEntity{
			CreatedAt: quote.BaseModel.CreatedAt,
			UpdatedAt: quote.BaseModel.UpdatedAt,
		},
	}, nil
}

func (q *QuotesRepo) UpdateQuote(quote entities.Quote) (entities.Quote, error) {
	panic("implement me")

}
