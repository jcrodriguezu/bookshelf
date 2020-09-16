package forms

import (
	"bookshelf/models"
)

// ReviewForm ...
type ReviewForm struct {
	Id     int    `form:",hidden, " valid:"int"`
	Title  string `form:"title" valid:"required"`
	Body   string `form:"body" valid:"required"`
	BookId int    `form:",hidden, " valid:"int"`
	UserId int    `form:",hidden, " valid:"int"`
}

// ToModel ...
func (f *ReviewForm) ToModel() (*models.Review, error) {
	book := &models.Book{Id: f.BookId}
	if err := book.Read(); err != nil {
		return nil, err
	}

	user := &models.User{Id: f.UserId}
	if err := user.Read(); err != nil {
		return nil, err
	}

	review := &models.Review{
		Id:    f.Id,
		Title: f.Title,
		Body:  f.Body,
		Book:  book,
		User:  user,
	}

	return review, nil
}
