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
	res := response.New("server is running")
	response.WriteJSON(a.Logger, w, res, http.StatusOK)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not create user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	passwordHash, err := generateHash(user.Password)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not create user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusInternalServerError)
		return
	}
	user.PasswordHash = passwordHash

	err = a.UserService.Create(r.Context(), &user)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not create user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	res := response.New("user created successfully").WithUser(user)
	response.WriteJSON(a.Logger, w, res, http.StatusCreated)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not update user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not update user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}
	user.ID = id

	passwordHash, err := generateHash(user.Password)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not update user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusInternalServerError)
		return
	}
	user.PasswordHash = passwordHash

	updates := make(map[string]any)
	updates["first_name"] = user.FirstName
	updates["last_name"] = user.LastName
	updates["email"] = user.Email
	updates["password"] = user.PasswordHash

	user, err = a.UserService.Update(r.Context(), id, updates)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not update user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	res := response.New("user updated successfully").WithUser(user)
	response.WriteJSON(a.Logger, w, res, http.StatusOK)
}

func (a *App) findUserByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not find user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	user, err := a.UserService.FindByID(r.Context(), id)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not find user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	res := response.New("found 1 user").WithUser(user)
	response.WriteJSON(a.Logger, w, res, http.StatusOK)
}

func (a *App) findAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.UserService.FindAll(r.Context())
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not retrieve users").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusInternalServerError)
		return
	}

	description := fmt.Sprintf("found %d users", len(users))
	res := response.New(description).WithUsers(users)
	response.WriteJSON(a.Logger, w, res, 200)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not delete user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	err = a.UserService.Delete(r.Context(), id)
	if err != nil {
		a.Logger.Error(err.Error())
		res := response.New("could not delete user").WithError(err)
		response.WriteJSON(a.Logger, w, res, http.StatusBadRequest)
		return
	}

	descripion := fmt.Sprintf("user with id %s was deleted successfully", idString)
	res := response.New(descripion)

	response.WriteJSON(a.Logger, w, res, http.StatusOK)
}
