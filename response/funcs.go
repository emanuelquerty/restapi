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

// Domain User includes password and password hash
// We do not want to include either in the response
// so we use response.User instead
func domainToResponseUser(user domain.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
