package main

import (
	"fmt"
	"quote/api/app/api/health"
	"quote/api/app/config"
	"quote/api/app/server"
	"quote/api/app/server/middleware"
	"quote/api/app/server/router"
	"quote/api/app/tools"
	"quote/api/app/tools/logger"

	"github.com/joho/godotenv"
	urlApi "github.com/sanctumlabs/curtz/api/url"
	"github.com/sanctumlabs/curtz/internal/repositories"
	"github.com/sanctumlabs/curtz/internal/services/urlsvc"

	"strconv"
)

const (
	Env                 = "ENV"
	EnvLogLevel         = "LOG_LEVEL"
	EnvLogJsonOutput    = "LOG_JSON_OUTPUT"
	EnvPort             = "PORT"
	EnvDatabaseHost     = "DATABASE_HOST"
	EnvDatabase         = "DATABASE"
	EnvDatabaseUsername = "DATABASE_USERNAME"
	EnvDatabasePassword = "DATABASE_PASSWORD"
	EnvDatabasePort     = "DATABASE_PORT"
)

func main() {
	log := logger.NewLogger("vehicle-api")

	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file. Using defaults")
	}

	environment := tools.EnvOr(Env, "development")
	logLevel := tools.EnvOr(EnvLogLevel, "debug")
	logJsonOutput := tools.EnvOr(EnvLogJsonOutput, "true")
	port := tools.EnvOr(EnvPort, "8080")
	host := tools.EnvOr(EnvDatabaseHost, "localhost")
	database := tools.EnvOr(EnvDatabase, "curtz-db")
	databaseUser := tools.EnvOr(EnvDatabaseUsername, "curtz-user")
	databasePass := tools.EnvOr(EnvDatabasePassword, "curtz-pass")
	databasePort := tools.EnvOr(EnvDatabasePort, "5432")

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
		Database: config.DatabaseConfig{
			Host:     host,
			Database: database,
			User:     databaseUser,
			Password: databasePass,
			Port:     databasePort,
		},
	}

	srv := server.NewServer(&configuration)

	// middlewares for the server
	corsMiddleware := middleware.NewCORSMiddleware(configuration.CorsHeaders)
	loggingMiddleware := middleware.NewLoggingMiddleware(configuration.Logging)
	recoveryMiddleware := middleware.NewRecoveryMiddleware()

	repository := repositories.NewRepository(configuration.Database)
	urlService := urlsvc.NewUrlService(repository.GetUrlRepo())

	// setup routers
	routers := []router.Router{
		urlApi.NewUrlRouter(urlService),
		health.NewHealthRouter(),
	}

	// initialize routers
	srv.InitRouter(routers...)

	// use middlewares
	srv.UseMiddleware(loggingMiddleware)
	srv.UseMiddleware(corsMiddleware)
	srv.UseMiddleware(recoveryMiddleware)

	appServer := srv.CreateServer()

	// start & run the server
	err = appServer.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		_, msg := fmt.Printf("Failed to start Server %s", err)
		log.Error(msg)
		panic(msg)
	}
}
