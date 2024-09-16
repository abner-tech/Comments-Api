package main

import (
	"encoding/json"
	"net/http"
)

// func (a *applicationDependences) healthChechHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, "status: available")
// 	// fmt.Fprintf(w, "environment: %s\n", a.config.environment)
// 	// fmt.Fprintf(w, "version: %s\n", appVersion)

// 	jsResponse := `{"Status": "available","environment": %q, "Version":%q}`
// 	jsResponse = fmt.Sprintf(jsResponse, a.config.environment, appVersion)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(jsResponse))
// }

func (a *applicationDependences) healthChechHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": a.config.environment,
		"version":     appVersion,
	}

	jsResponse, err := json.Marshal(data)
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w,
			"The Server Encountered a problem and could not proccess your Request",
			http.StatusInternalServerError)
		return
	}

	jsResponse = append(jsResponse, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsResponse)
}
