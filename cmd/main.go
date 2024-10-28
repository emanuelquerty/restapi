package main

import (
	"log"
	"log/slog"
	"os"
	"restapi/app"
	"restapi/middleware"
	"restapi/storage/sqlite"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DSN  string
}

func loadConfigFromEnv(config *Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	dsn := os.Getenv("DB_DSN")

	config.Port = port
	config.DSN = dsn
}

func main() {
	config := &Config{}
	loadConfigFromEnv(config)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := sqlite.NewDB(config.DSN)
	if err != nil {
		logger.Error("could not open database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	app := app.New(logger, db)

	app.Use(middleware.AccessLogger)
	app.Use(middleware.SecurityHeaders)

	err = app.ListenAndServe(config.Port)
	if err != nil {
		logger.Error("could not start server", slog.String("error", err.Error()))
	}
}
