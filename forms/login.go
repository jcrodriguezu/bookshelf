package forms

import (
	"bookshelf/models"
)

// LoginForm ...
type LoginForm struct {
	Username string `form:"username" valid:"required"`
	Password string `form:"password,password,Password: " valid:"required"`
}

// ToModel ...
func (f *LoginForm) ToModel() (models.IModel, error) {
	user := &models.User{
		Username: f.Username,
		Password: f.Password,
	}
	return user, nil
}
