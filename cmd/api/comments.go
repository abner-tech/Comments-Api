package main

import (
	"fmt"
	"net/http"
	//import data package whichcontains the definition for Comment
	//"github.com/abner-tech/Comments-Api/internal/data"
)

func (a *applicationDependences) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	//create a struct to hold a comment
	//we use struct tags [` `] to make the names display in lowercase

	var incomingData struct {
		Content string `json: "Content"`
		Author  string `json:"author"`
	}

	//perform decoding

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	//for now display the result
	fmt.Fprintf(w, "%+v\n", incomingData)
}
