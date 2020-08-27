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

// GetByUserNameAndPassword ...
func (u *User) GetByUserNameAndPassword(username string, password string) (*User, error) {
	var user User
	qs := orm.NewOrm().QueryTable(new(User))
	err := qs.Filter("Username", username).Filter("Password", password).One(&user)
	if err == orm.ErrNoRows {
		return nil, fmt.Errorf("User not found")
	}

	return &user, nil
}
