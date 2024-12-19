package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/domain"
	"restapi/response"
	"strconv"
)

func (a *App) checkHealth(w http.ResponseWriter, r *http.Request) *appError {
	res := struct {
		Message string
	}{"Server is running"}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "Invalid json body", http.StatusInternalServerError}
	}
	return nil
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) *appError {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return &appError{err, "Invalid json body", http.StatusBadRequest}
	}

	err = user.HashPassword()
	if err != nil {
		return &appError{err, "Could not hash password", http.StatusInternalServerError}
	}

	err = a.UserService.Create(r.Context(), &user)
	if err != nil {
		return &appError{err, "Could not create user", http.StatusInternalServerError}
	}

	w.WriteHeader(http.StatusCreated)
	res := response.New("User created successfully").WithUser(user)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "Internal error", http.StatusInternalServerError}
	}
	return nil
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) *appError {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return &appError{err, "internal error", http.StatusInternalServerError}
	}

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return &appError{err, "could not update user", http.StatusInternalServerError}
	}
	user.ID = id

	err = user.HashPassword()
	if err != nil {
		return &appError{err, "Could not hash password", http.StatusInternalServerError}
	}

	updates := make(map[string]any)
	updates["first_name"] = user.FirstName
	updates["last_name"] = user.LastName
	updates["email"] = user.Email
	updates["password"] = user.PasswordHash

	user, err = a.UserService.Update(r.Context(), id, updates)
	if err != nil {
		return &appError{err, "could not update user", http.StatusInternalServerError}
	}

	w.WriteHeader(http.StatusCreated)
	res := response.New("user updated successfully").WithUser(user)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "internal error", http.StatusCreated}
	}
	return nil
}

func (a *App) findUserByID(w http.ResponseWriter, r *http.Request) *appError {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return &appError{err, "bad request", http.StatusBadRequest}
	}

	user, err := a.UserService.FindByID(r.Context(), id)
	if err != nil {
		return &appError{err, "could not find user", http.StatusBadRequest}
	}

	w.WriteHeader(http.StatusOK)
	res := response.New("found 1 user").WithUser(user)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "internal error", http.StatusCreated}
	}
	return nil
}

func (a *App) findAllUsers(w http.ResponseWriter, r *http.Request) *appError {
	users, err := a.UserService.FindAll(r.Context())
	if err != nil {
		return &appError{err, "could not retrieve users", http.StatusInternalServerError}
	}

	description := fmt.Sprintf("found %d users", len(users))
	res := response.New(description).WithUsers(users)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "internal error", http.StatusCreated}
	}
	return nil
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) *appError {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return &appError{err, "bad request", http.StatusInternalServerError}
	}

	err = a.UserService.Delete(r.Context(), id)
	if err != nil {
		return &appError{err, "could not delete user", http.StatusBadRequest}
	}

	descripion := fmt.Sprintf("user with id %s was deleted successfully", idString)
	res := response.New(descripion)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return &appError{err, "internal error", http.StatusCreated}
	}
	return nil
}
