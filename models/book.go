package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Book model
type Book struct {
	Id      int `orm:"pk"`
	Title   string
	Author  string
	Copies  int       `orm:"default(0)"`
	Users   []*User   `orm:"reverse(many)"`
	Reviews []*Review `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Book))
}

// AvailableCopies number.
func (b *Book) AvailableCopies() int {
	return b.Copies - len(b.Users)
}

// All returns the list of all books.
func (b *Book) All() []*Book {
	var books []*Book
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Book)).All(&books)
	if err != nil {
		fmt.Println("Error getting all books")
		return nil
	}

	return books
}
