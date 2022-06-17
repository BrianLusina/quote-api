package repositories

import (
	"fmt"
	"log"
	"os"
	"quote/api/app/config"
	"quote/api/app/internal/repositories/models"
	"quote/api/app/internal/repositories/quotesrepo"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"
)

type repository struct {
	db         *gorm.DB
	quotesRepo *quotesrepo.QuotesRepo
}

func NewRepository(config config.DatabaseConfig) *repository {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.Database)

	dbLogger := dblogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		dblogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  dblogger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		log.Fatalf("DB Connection failed with err: %v", err)
	}

	if err = db.AutoMigrate(&models.Quote{}); err != nil {
		log.Fatalf("AutoMigration failed with err: %v", err)
	}

	return &repository{
		db:         db,
		quotesRepo: quotesrepo.NewQuotesRepo(db),
	}
}

func (r repository) GetQuotesRepo() *quotesrepo.QuotesRepo {
	return r.quotesRepo
}
