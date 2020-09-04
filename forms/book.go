package forms

import (
	"bookshelf/models"
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// BookForm ...
type BookForm struct {
	Id     int    `form:",hidden, " valid:"int"`
	Title  string `form:"title" valid:"required"`
	Author string `form:"author" valid:"required"`
	Copies int    `form:"copies" valid:"int, required"`
}

// ToModel ...
func (f *BookForm) ToModel() (*models.Book, error) {
	isValid, err := valid.ValidateStruct(f)
	if !isValid {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("All the fields are required")
	}

	book := &models.Book{
		Id:     f.Id,
		Title:  f.Title,
		Author: f.Author,
		Copies: f.Copies,
	}

	return book, nil
}
