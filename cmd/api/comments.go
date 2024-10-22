package main

import (
	"fmt"
	"net/http"

	"github.com/abner-tech/Comments-Api.git/internal/data"
	"github.com/abner-tech/Comments-Api.git/internal/validator"
)

func (a *applicationDependences) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	//create a struct to hold a comment
	//we use struct tags [` `] to make the names display in lowercase
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	//perform decoding

	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	comment := &data.Comment{
		Content: incomingData.Content,
		Author:  incomingData.Author,
	}

	v := validator.New()
	//do validation
	data.ValidateComment(v, comment)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors) //implemented later
		return
	}

	//add comment to the comments table in database
	err = a.commentModel.Insert(comment)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	//for now display the result
	// fmt.Fprintf(w, "%+v\n", incomingData)

	//set a location header, the path to the newly created comments
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/comments/%d", comment.ID))

	//send a json response with a 201 (new reseource created) status code
	data := envelope{
		"comment": comment,
	}
	err = a.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
