package restapi

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"restapi/app"
	"restapi/domain"
	"restapi/middleware"
	"restapi/response"
	"restapi/storage/sqlite"
	"testing"
)

var Users = []domain.User{
	{
		FirstName: "Peter",
		LastName:  "Petrelli",
		Email:     "ppetrelli@email.com",
		Password:  "soME2050PASS",
	},
	{
		FirstName: "Leroy",
		LastName:  "Jenkis",
		Email:     "ljenkins@email.com",
		Password:  "62hsROOMmie273ms",
	},
}

func TestApp(t *testing.T) {
	go func() {
		setupAndStartServer()
	}()

	client := &http.Client{}

	// Create two users
	bodyBytes, _ := json.Marshal(Users[0])
	buf := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/users", buf)
	if err != nil {
		t.Fatalf("error creating user: %s", err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("error sending request: %s", err)
	}

	var serverResponse response.Response
	json.NewDecoder(res.Body).Decode(&serverResponse)

}

func setupAndStartServer() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := sqlite.NewDB("./api_test.db")
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
