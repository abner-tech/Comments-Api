package data

import (
	"time"
)

// each name begins with uppercase to make them exportable/public
type Comment struct {
	ID        int64     //unique value per comment
	Content   string    //comment data
	Author    string    //person who wrote comment
	CreatedAt time.Time //database timestamp
	Version   int32     //icremented on each update
}
