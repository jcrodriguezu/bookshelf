package forms

import (
	"bookshelf/models"
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// LoginForm ...
type LoginForm struct {
	Username interface{} `form:"username" valid:"required"`
	Password interface{} `form:"password" valid:"required"`
}

// GetData ...
func (f *LoginForm) GetData() (*models.User, error) {
	isValid, _ := valid.ValidateStruct(f)
	if !isValid {
		return nil, fmt.Errorf("Username and Password can't be empty")
	}

	user := &models.User{
		Username: f.Username.(string),
		Password: f.Password.(string),
	}
	return user, nil
}
