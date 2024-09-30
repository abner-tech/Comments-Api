package main

import (
	"net/http"
)

func (a *applicationDependences) healthChechHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.config.environment,
			"version":     appVersion,
		},
	}

	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, "the server encountered a problem and cound not process your request", http.StatusInternalServerError)
	}
}
