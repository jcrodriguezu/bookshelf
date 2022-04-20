package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
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

// AllByBook ...
func (b *Review) AllByBook(bookId int) []*Review {
	var reviews []*Review
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Review)).Filter("Book__Id", bookId).All(&reviews)
	if err != nil {
		fmt.Println("Error getting all books")
		return nil
	}

	return reviews
}

// Insert ...
func (b *Review) Insert() error {
	o := orm.NewOrm()

	if _, err := o.Insert(b); err != nil {
		return err
	}

	return nil
}
