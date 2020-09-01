package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// User model.
type User struct {
	Id        int       `orm:"pk, auto"`
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
func (u *User) DoLogin() error {
	pass := u.Password
	u.Password = ""

	o := orm.NewOrm()
	err := o.Read(u, "Username")

	// TODO user.Password shouldn't be un plain text
	if err == nil && u.Password == pass {
		if _, err := o.LoadRelated(u, "Role"); err != nil {
			return fmt.Errorf("Error loading role for user")
		}
		return nil
	}

	return fmt.Errorf("Wrong Username or Password")
}
