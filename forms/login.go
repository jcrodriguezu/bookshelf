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

// GetData ...
func (f *LoginForm) GetData() (*models.User, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf("Username and Password can't be empty")
	}

	u := new(models.User)
	u.Username = f.Username.(string)
	u.Password = f.Password.(string)

	return u, nil
}
