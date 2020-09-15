package controllers

import (
	"bookshelf/models"

	"github.com/astaxie/beego"
)

// UserBookController ...
type UserBookController struct {
	beego.Controller
}

// LendBook ...
func (c *UserBookController) LendBook() {
	user := c.GetSession("user")

	flash := beego.NewFlash()

	bookId, err := c.GetInt("bookid")
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/index", 307)
	}

	if err = user.(*models.User).LendBook(bookId); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	c.Redirect("/index", 307)
}

// ReturnBook ...
func (c *UserBookController) ReturnBook() {
	user := c.GetSession("user")

	flash := beego.NewFlash()

	bookId, err := c.GetInt("bookid")
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/index", 307)
	}

	if err = user.(*models.User).ReturnBook(bookId); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	c.Redirect("/index", 307)
}
