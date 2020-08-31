package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// User model.
type User struct {
	Id        int       `orm:"pk"`
	Name      string    `orm:"unique"`
	Username  string    `orm:"unique"`
	Password  string    `orm:"unique"`
	Role      *Role     `orm:"null;rel(one);on_delete(set_null)"`
	BooksLent []*Book   `orm:"rel(m2m)"`
	Reviews   []*Review `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// DoLogin ...
func (u *User) DoLogin(username string, password string) (*User, error) {
	toCheck := &User{Username: username}

	o := orm.NewOrm()
	err := o.Read(toCheck, "Username")

	// TODO user.Password shouldn't be un plain text
	if err == nil && password == toCheck.Password {
		if _, err := o.LoadRelated(toCheck, "Role"); err != nil {
			return nil, fmt.Errorf("Error loading role for user")
		}
		return toCheck, nil
	}

	return nil, fmt.Errorf("Wrong Username or Password")
}
