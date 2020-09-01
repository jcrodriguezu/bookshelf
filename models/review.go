package models

import (
	"github.com/astaxie/beego/orm"
)

// Review model
type Review struct {
	Id    int `orm:"pk, auto"`
	Title string
	Body  string
	Book  *Book `orm:"rel(fk)"`
	User  *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Review))
}
