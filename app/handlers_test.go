package app_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"restapi/app"
	"restapi/domain"
	"restapi/response"
	"restapi/storage/inmemory"
	"restapi/testdata"
	"testing"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	users  = testdata.Users

	userStore = inmemory.NewUserStore(users)
)

func TestUserEndpoints(t *testing.T) {
	app := app.New(logger, nil)
	app.UserService = userStore
	app.RegisterRoutes()

	t.Run("find user by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users/1", nil)
		resp := httptest.NewRecorder()

		app.Mux.ServeHTTP(resp, req)

		wantUser := response.User{
			ID:        1,
			FirstName: "Peter",
			LastName:  "Petrelli",
			Email:     "ppetrelli@email.com",
		}

		var gotResp response.Response

		err := json.NewDecoder(resp.Body).Decode(&gotResp)
		if err != nil {
			t.Fatalf("got %s but want nil", err)
		}

		if !reflect.DeepEqual(gotResp.User, wantUser) {
			t.Fatalf("got %+v but want %+v", gotResp.User, wantUser)
		}

	})

	t.Run("creates and returns new user", func(t *testing.T) {
		user := domain.User{
			FirstName: "Freddy",
			LastName:  "Krueger",
			Email:     "freddywillfindyou@email.com",
			Password:  "Nightmare In Elm Street",
		}

		userBytes, _ := json.Marshal(user)
		buf := bytes.NewBuffer(userBytes)

		req := httptest.NewRequest(http.MethodPost, "/api/users", buf)
		resp := httptest.NewRecorder()

		app.Mux.ServeHTTP(resp, req)

		// id of new user should be 3
		wantUser := response.User{
			ID:        3,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}

		var gotResp response.Response

		err := json.NewDecoder(resp.Body).Decode(&gotResp)
		if err != nil {
			t.Fatalf("got %s but want nil", err)
		}

		if !reflect.DeepEqual(gotResp.User, wantUser) {
			t.Fatalf("got %+v but want %+v", gotResp.User, wantUser)
		}

	})
}
