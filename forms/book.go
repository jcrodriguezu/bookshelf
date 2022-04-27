package forms

import (
	"bookshelf/models"

	"github.com/asaskevich/govalidator"
)

// BookForm ...
type BookForm struct {
	Id     int    `form:",hidden, " valid:"int"`
	Isbn   string `form:"Isbn" id:"isbn" valid:"isbn"`
	Title  string `form:"title" id:"title" valid:"required"`
	Author string `form:"author" id:"author" valid:"required"`
	Copies int    `form:"copies" valid:"int, required"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("isbn", func(i interface{}, o interface{}) bool {
		if govalidator.IsISBN10(i.(string)) || govalidator.IsISBN13(i.(string)) {
			return true
		}
		return false
	})
}

// ToModel ...
func (f *BookForm) ToModel() (models.IModel, error) {
	book := &models.Book{
		Id:     f.Id,
		Isbn:   f.Isbn,
		Title:  f.Title,
		Author: f.Author,
		Copies: f.Copies,
	}

	return book, nil
}
