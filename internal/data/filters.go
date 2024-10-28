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

type Metadata struct {
	CurrentPage  int `json:"current_page, omitempty"`
	PageSize     int `json:"page_size, omitempty"`
	FirstPage    int `json:"page_size, omitempty"`
	LastPage     int `json:"last_page, omitempty"`
	TotalRecords int `json:"total_records, omitempty"`
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

// Calculate the matadata
func calculateMetaData(totalRecords int, currentPage int, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  currentPage,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}
