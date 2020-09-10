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

func (u *User) Read() error {
	o := orm.NewOrm()
	if err := o.Read(u); err != nil {
		return err
	}

	if _, err := o.LoadRelated(u, "BooksLent"); err != nil {
		return err
	}

	if _, err := o.LoadRelated(u, "Role"); err != nil {
		return err
	}

	return nil
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

// LendBook ...
func (u *User) LendBook(bookid int) error {
	book := &Book{Id: bookid}
	o := orm.NewOrm()
	if o.QueryM2M(u, "BooksLent").Exist(book) {
		return fmt.Errorf("User already has the book")
	}

	if err := o.Read(book); err != nil {
		return err
	}

	if book.AvailableCopies() <= 0 {
		return fmt.Errorf("The book has no copies available")
	}

	if _, err := o.QueryM2M(u, "BooksLent").Add(book); err != nil {
		return err
	}

	return nil
}

// ReturnBook ...
func (u *User) ReturnBook(bookid int) error {
	book := &Book{Id: bookid}

	o := orm.NewOrm()
	if !o.QueryM2M(u, "BooksLent").Exist(book) {
		return fmt.Errorf("User doesn't has the book")
	}

	if err := o.Read(book); err != nil {
		return err
	}

	if _, err := o.QueryM2M(u, "BooksLent").Remove(book); err != nil {
		return err
	}

	return nil
}

// HasBook ...
func HasBook(userid int, bookid int) bool {
	book := &Book{Id: bookid}
	user := &User{Id: userid}
	o := orm.NewOrm()
	if err := o.Read(user); err != nil {
		fmt.Println(err)
		return false
	}
	return o.QueryM2M(user, "BooksLent").Exist(book)
}
