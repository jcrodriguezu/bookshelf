package forms

import (
	"bookshelf/models"
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// LoginForm ...
type LoginForm struct {
	Username string `form:"username" valid:"required"`
	Password string `form:"password" valid:"required"`
}

// ToModel ...
func (f *LoginForm) ToModel() (*models.User, error) {
	isValid, _ := valid.ValidateStruct(f)
	if !isValid {
		return nil, fmt.Errorf("Username and Password can't be empty")
	}

	user := &models.User{
		Username: f.Username,
		Password: f.Password,
	}
	return user, nil
}
