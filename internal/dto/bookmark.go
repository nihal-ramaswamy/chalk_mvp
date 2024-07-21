package dto

type Bookmark struct {
	StudentEmail string `json:"student_email"`
}

func NewBookmark() *Bookmark {
	return &Bookmark{}
}
