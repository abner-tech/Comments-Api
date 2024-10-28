package data

import (
	"github.com/abner-tech/Comments-Api.git/internal/validator"
)

//the Filters type will contain the fields related to pagination
//and eventually the fields related to sorting

type Fileters struct {
	Page     int //which page number does the client want
	PageSize int //how many records per page
}

//we validate page and Page size
//follow same approach used to validate a comment

func ValidateFilters(v *validator.Validator, f Fileters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 500, "page", "must be a maximim of 500")
	v.Check(f.PageSize > 0, "page_size", "must be greator than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximun of 100")
}

// calculate how many results to send back
func (f Fileters) limit() int {
	return f.PageSize
}

// calculate the offset so that we remember how many records have been sent
// and how many remain to be sent
func (f Fileters) offset() int {
	return (f.Page - 1) * f.PageSize
}
