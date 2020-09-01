package forms

import (
	"bookshelf/models"
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// BookForm ...
type BookForm struct {
	Title  interface{} `form:"title" valid:"required"`
	Author interface{} `form:"author" valid:"required"`
	Copies interface{} `form:"copies" valid:"int, required"`
}

// GetData ...
func (f *BookForm) GetData() (*models.Book, error) {
	isValid, err := valid.ValidateStruct(f)
	if !isValid {
		return nil, fmt.Errorf("All the fields are required")
	}

	copies, err := valid.ToInt(f.Copies)
	if err != nil {
		return nil, fmt.Errorf("Number of copies should be a number greater than 0")
	}

	book := &models.Book{
		Title:  f.Title.(string),
		Author: f.Author.(string),
		Copies: int(copies),
	}

	return book, nil
}
