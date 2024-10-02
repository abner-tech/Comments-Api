package main

import (
	"fmt"
	"net/http"
)

func (a *applicationDependences) logError(r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()
	a.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (a *applicationDependences) errorResponseJSON(w http.ResponseWriter, r *http.Request, status int, message any) {
	errorData := envelope{"error": message}
	err := a.writeJSON(w, status, errorData, nil)
	if err != nil {
		a.logError(r, err)
		w.WriteHeader(500)
	}
}

// send an error response if our server messes up
func (a *applicationDependences) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//first thing is to log error message
	a.logError(r, err)
	//prepare a response to send to the client
	message := "the server encountered a problem and cound not process your request"
	a.errorResponseJSON(w, r, http.StatusInternalServerError, message)
}

// send an error response of our client messes up with a 404
func (a *applicationDependences) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	//we only log server errors, not client errors
	//prepare a response to send to the client
	message := "the requested resource cound not be found"
	a.errorResponseJSON(w, r, http.StatusNotFound, message)
}

// semd an error response if our client messes up with 405
func (a *applicationDependences) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	//we only log server errors, not client errors
	//prepare a FORMATED response to send to the client
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	a.errorResponseJSON(w, r, http.StatusMethodNotAllowed, message)
}