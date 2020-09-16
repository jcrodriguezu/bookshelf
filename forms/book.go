package forms

import (
	"bookshelf/models"
)

// BookForm ...
type BookForm struct {
	Id     int    `form:",hidden, " valid:"int"`
	Title  string `form:"title" valid:"required"`
	Author string `form:"author" valid:"required"`
	Copies int    `form:"copies" valid:"int, required"`
}

// ToModel ...
func (f *BookForm) ToModel() (models.IModel, error) {
	book := &models.Book{
		Id:     f.Id,
		Title:  f.Title,
		Author: f.Author,
		Copies: f.Copies,
	}

	return book, nil
}
