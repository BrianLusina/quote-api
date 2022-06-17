package main

import (
	"context"
	"fmt"
	"quote/api/app/api/health"
	"quote/api/app/api/quotes"
	"quote/api/app/config"
	"quote/api/app/internal/core/services/quotesvc"
	"quote/api/app/internal/repositories"
	"quote/api/app/server"
	"quote/api/app/server/middleware"
	"quote/api/app/server/router"
	"quote/api/app/tools"
	"quote/api/app/tools/logger"
	"quote/api/app/utils/cache"
	"quote/api/app/utils/monitoring"

	"github.com/joho/godotenv"

	"strconv"
)

const (
	Env                 = "ENV"
	EnvLogLevel         = "LOG_LEVEL"
	EnvLogJsonOutput    = "LOG_JSON_OUTPUT"
	EnvPort             = "PORT"
	EnvDatabaseHost     = "DATABASE_HOST"
	EnvDatabase         = "DATABASE_NAME"
	EnvDatabaseUsername = "DATABASE_USERNAME"
	EnvDatabasePassword = "DATABASE_PASSWORD"
	EnvDatabasePort     = "DATABASE_PORT"
	EnvDatabaseSSLMode  = "DATABASE_SSLMODE"
	EnvSentryDsn        = "SENTRY_DSN"
	EnvCacheHost        = "CACHE_HOST"
	EnvCachePort        = "CACHE_PORT"
	EnvCacheDb          = "CACHE_DB"
	EnvCacheUsername    = "CACHE_USERNAME"
	EnvCachePassword    = "CACHE_PASSWORD"
	EnvUsername         = "USERNAME"
	EnvPassword         = "PASSWORD"
)

func main() {
	ctx := context.Background()

	log := logger.NewLogger("quotes-api")

	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file. Using defaults")
	}

	environment := tools.EnvOr(Env, "development")
	logLevel := tools.EnvOr(EnvLogLevel, "debug")
	logJsonOutput := tools.EnvOr(EnvLogJsonOutput, "true")
	port := tools.EnvOr(EnvPort, "8080")
	host := tools.EnvOr(EnvDatabaseHost, "localhost")
	database := tools.EnvOr(EnvDatabase, "quotesdb")
	databaseUser := tools.EnvOr(EnvDatabaseUsername, "quotesUser")
	databasePass := tools.EnvOr(EnvDatabasePassword, "quotesPass")
	databasePort := tools.EnvOr(EnvDatabasePort, "5432")
	databaseSslMode := tools.EnvOr(EnvDatabaseSSLMode, "disable")
	sentryDsn := tools.EnvOr(EnvSentryDsn, "")
	cacheHost := tools.EnvOr(EnvCacheHost, "localhost")
	cacheDb := tools.EnvOr(EnvCacheDb, "0")
	cachePort := tools.EnvOr(EnvCachePort, "6379")
	cacheUsername := tools.EnvOr(EnvCacheUsername, "quotesUser")
	cachePassword := tools.EnvOr(EnvCachePassword, "quotesPass")
	username := tools.EnvOr(EnvUsername, "admin")
	password := tools.EnvOr(EnvPassword, "admin")
	allowedOrigins := tools.EnvOr("ALLOWED_ORIGINS", "*")

	enableJsonOutput, err := strconv.ParseBool(logJsonOutput)
	if err != nil {
		enableJsonOutput = true
	}

	configuration := config.Config{
		Env:  environment,
		Port: port,
		Logging: config.LoggingConfig{
			Level:            logLevel,
			EnableJSONOutput: enableJsonOutput,
		},
		Cors: config.CorsConfig{
			AllowedOrigins: allowedOrigins,
		},
		Database: config.DatabaseConfig{
			Host:     host,
			Database: database,
			User:     databaseUser,
			Password: databasePass,
			Port:     databasePort,
			SSLMode:  databaseSslMode,
		},
		Monitoring: config.Monitoring{
			Sentry: config.Sentry{
				Dsn: sentryDsn,
			},
		},
		Cache: config.CacheConfig{
			Host:     cacheHost,
			Port:     cachePort,
			Db:       cacheDb,
			Username: cacheUsername,
			Password: cachePassword,
		},
		Auth: config.AuthConfig{
			Username: username,
			Password: password,
		},
	}

	cache := cache.New(ctx, configuration.Cache)

	monitoring.InitializeMonitoring(configuration.Monitoring)

	srv := server.NewServer(&configuration)

	// middlewares for the server
	corsMiddleware := middleware.NewCORSMiddleware(configuration.Cors)
	loggingMiddleware := middleware.NewLoggingMiddleware(configuration.Logging)
	recoveryMiddleware := middleware.NewRecoveryMiddleware()
	monitoringMiddleware := middleware.NewMonitoringMiddleware()
	authMiddleware := middleware.NewAuthenticationMiddleware(configuration.Auth)

	// use middlewares
	srv.UseMiddleware(recoveryMiddleware)
	srv.UseMiddleware(monitoringMiddleware)
	srv.UseMiddleware(loggingMiddleware)
	srv.UseMiddleware(corsMiddleware)
	srv.UseMiddleware(authMiddleware)

	repository := repositories.NewRepository(configuration.Database)
	quotesService := quotesvc.NewQuoteSvc(repository.GetQuotesRepo())

	// setup routers
	routers := []router.Router{
		quotes.NewQuotesRouter(cache, quotesService),
		health.NewHealthRouter(),
	}

	// initialize routers
	srv.InitRouter(routers...)

	appServer := srv.CreateServer()

	// start & run the server
	err = appServer.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		_, msg := fmt.Printf("Failed to start Server %s", err)
		log.Error(msg)
		panic(msg)
	}
}
