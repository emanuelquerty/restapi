package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"restapi/domain"
)

func WriteJSON(logger *slog.Logger, w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error("error encoding json", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MapToUser(user domain.User) UserWithNoPasswordField {
	var resUser UserWithNoPasswordField

	resUser.ID = user.ID
	resUser.FirstName = user.FirstName
	resUser.LastName = user.LastName
	resUser.Email = user.Email

	return resUser
}

func MapToUsers(users []domain.User) []UserWithNoPasswordField {
	var resUsers []UserWithNoPasswordField
	for _, user := range users {
		var currUser UserWithNoPasswordField

		currUser.ID = user.ID
		currUser.FirstName = user.FirstName
		currUser.LastName = user.LastName
		currUser.Email = user.Email

		resUsers = append(resUsers, currUser)
	}
	return resUsers
}
