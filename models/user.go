package models

import (
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

// GetByUsername ...
func GetByUsername(username string) *User {
	user := User{Username: username}
	orm.NewOrm().Read(&user, "Username")
	return &user
}
