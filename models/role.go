package models

import (
	"github.com/astaxie/beego/orm"
)

// Role model
type Role struct {
	Id   int    `orm:"pk"`
	Name string `orm:"unique"`
	User *User  `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Role))
}
