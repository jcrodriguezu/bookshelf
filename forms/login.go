package forms

import (
	"bookshelf/models"
	"fmt"
)

// LoginForm ...
type LoginForm struct {
	Username interface{} `form:"username"`
	Password interface{} `form:"password"`
}

// IsValid ...
func (f *LoginForm) IsValid() bool {
	return f.Username != nil && f.Password != nil
}

// DoLogin ...
func (f *LoginForm) DoLogin() (*models.User, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("Username and Password can't be empty")
	}

	user := models.GetByUsername(f.Username.(string))
	// TODO user.Password shouldn't be un plain text
	if user != nil && user.Password == f.Password.(string) {
		return user, nil
	}

	return nil, fmt.Errorf("Wrong Username or Password")
}
