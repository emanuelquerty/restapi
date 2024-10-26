package main

import (
	"log/slog"
	"os"
	"restapi/app"
	"restapi/middleware"
	"restapi/storage/sqlite"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := sqlite.NewDB("./api.db")
	if err != nil {
		logger.Error("could not open database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	app := app.New(logger, db)

	app.Use(middleware.AccessLogger)
	app.Use(middleware.SecurityHeaders)

	err = app.ListenAndServe("3000")
	if err != nil {
		logger.Error("could not start server", slog.String("error", err.Error()))
	}
}
