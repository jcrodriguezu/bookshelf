package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

// Book model
type Book struct {
	Id      int    `orm:"pk, auto"`
	Isbn    string `orm:"unique"`
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
	o := orm.NewOrm()
	if _, err := o.LoadRelated(b, "Users"); err != nil {
		return 0 // TODO should this return an error too?
	}
	return b.Copies - len(b.Users)
}

// GetReviews ...
func (b *Book) GetReviews() []*Review {
	o := orm.NewOrm()
	if _, err := o.LoadRelated(b, "Reviews"); err != nil {
		return nil
	}
	return b.Reviews
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

// GetById ...
func (b *Book) Read() error {
	o := orm.NewOrm()

	if err := o.Read(b); err != nil {
		return err
	}

	if _, err := o.LoadRelated(b, "Reviews"); err != nil {
		return err
	}

	return nil
}

// Insert ...
func (b *Book) Insert() error {
	o := orm.NewOrm()

	if _, err := o.Insert(b); err != nil {
		return err
	}

	return nil
}

// Update ...
func (b *Book) Update() error {
	o := orm.NewOrm()

	if _, err := o.Update(b); err != nil {
		return err
	}

	return nil
}

// Delete ...
func (b *Book) Delete() error {
	o := orm.NewOrm()

	if _, err := o.Delete(b); err != nil {
		return err
	}

	return nil
}
