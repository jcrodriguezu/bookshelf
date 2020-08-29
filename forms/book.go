package forms

import (
	"bookshelf/models"
	"fmt"
)

// BookForm ...
type BookForm struct {
	Title  interface{} `form:"title"`
	Author interface{} `form:"author"`
	Copies interface{} `form:"copies"`
}

// IsValid ...
func (f *BookForm) IsValid() bool {
	return f.Title != nil && f.Author != nil
}

// DoCreate ...
func (f *BookForm) DoCreate() (*models.Book, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("Title and Author can't be empty")
	}

	b := new(models.Book)
	b.Title = f.Title.(string)
	b.Author = f.Author.(string)
	b.Copies = f.Copies.(int)

	if err := b.Insert(); err != nil {
		return nil, err
	}

	return b, nil
}
