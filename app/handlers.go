package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/domain"
	"restapi/response"
	"strconv"
)

func (a *App) checkHealth(w http.ResponseWriter, r *http.Request) {
	res := struct {
		Msg string `json:"msg,omitempty"`
	}{
		Msg: "server is running",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = user.HashPassword()
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.userStore.Create(r.Context(), &user)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.SingleUser{
		Description: "user created successfully",
		User:        response.MapToUser(user),
	}

	response.WriteJSON(a.Logger, w, res, 201)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = id

	err = user.HashPassword()
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updates := make(map[string]any)
	updates["first_name"] = user.FirstName
	updates["last_name"] = user.LastName
	updates["email"] = user.Email
	updates["password"] = user.PasswordHash

	user, err = a.userStore.Update(r.Context(), id, updates)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.SingleUser{
		Description: "user updated successfully",
		User:        response.MapToUser(user),
	}

	response.WriteJSON(a.Logger, w, res, 200)
}

func (a *App) findUserByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.userStore.FindByID(r.Context(), id)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := response.SingleUser{
		Description: "found 1 user",
		User:        response.MapToUser(user),
	}

	response.WriteJSON(a.Logger, w, res, 200)
}

func (a *App) findAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.userStore.FindAll(r.Context())
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	descritption := fmt.Sprintf("found %d users", len(users))
	res := response.MultiUser{
		Description: descritption,
		Users:       response.MapToUsers(users),
	}

	response.WriteJSON(a.Logger, w, res, 200)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.userStore.Delete(r.Context(), id)
	if err != nil {
		a.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("user with id %s was deleted successfully", idString)
	res := response.Generic{Description: msg}

	w.Header().Set("Content-Type", "application/json")
	response.WriteJSON(a.Logger, w, res, http.StatusOK)
}
