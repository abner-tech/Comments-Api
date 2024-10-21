package data

import (
	"time"

	"github.com/abner-tech/Comments-Api.git/internal/validator"
)

// each name begins with uppercase to make them exportable/public
type Comment struct {
	ID        int64     `json:"id"`      //unique value per comment
	Content   string    `json:"content"` //comment data
	Author    string    `json:"author"`  //person who wrote comment
	CreatedAt time.Time `json:"-"`       //database timestamp
	Version   int32     `json:"version"` //icremented on each update
}

func ValidateComment(v *validator.Validator, comment *Comment) {
	//check if the content field is empty
	v.Check(comment.Content != "", "content", "must be provided")
	//check if the Author field is empty
	v.Check(comment.Author != "", "author", "must be provided")
	//check if the content field is empty
	v.Check(len(comment.Content) <= 100, "content", "must not be more than 100 bytes long")
	//check is author field is empty
	v.Check(len(comment.Author) <= 25, "author", "must not be more than 25 bytes long")
}
