package main

import (
	"errors"
	"net/http"

	"github.com/abner-tech/Comments-Api.git/internal/data"
	"github.com/abner-tech/Comments-Api.git/internal/validator"
)

func (a *applicationDependences) registerUserhandler(w http.ResponseWriter, r *http.Request) {

	var incomingData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
	}

	user := &data.User{
		Username:  incomingData.Username,
		Email:     incomingData.Email,
		Activated: false,
	}

	//hashing password and storing with the cleartect version
	err = user.Password.Set(incomingData.Password)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	//user valudation
	v := validator.New()
	data.ValidateUser(v, user)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.userModel.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	data := envelope{
		"user": user,
	}

	//status code 201 resource created
	err = a.writeJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
