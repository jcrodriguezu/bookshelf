package forms

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// LoginForm ...
type LoginForm struct {
	Username interface{} `form:"username" valid:"required"`
	Password interface{} `form:"password" valid:"required"`
}

// GetData ...
func (f *LoginForm) GetData() (string, string, error) {
	isValid, _ := valid.ValidateStruct(f)
	if !isValid {
		return "", "", fmt.Errorf("Username and Password can't be empty")
	}

	return f.Username.(string), f.Password.(string), nil
}
