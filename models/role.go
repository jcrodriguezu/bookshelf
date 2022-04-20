package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// Role model
type Role struct {
	Id   int    `orm:"pk, auto"`
	Name string `orm:"unique"`
	User *User  `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Role))
}
