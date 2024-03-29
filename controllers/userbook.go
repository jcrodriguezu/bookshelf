package controllers

import (
	"bookshelf/models"

	"github.com/beego/beego/v2/server/web"
)

// UserBookController ...
type UserBookController struct {
	web.Controller
}

// LendBook ...
func (c *UserBookController) LendBook() {
	user := c.GetSession("user")

	flash := web.NewFlash()

	bookId, err := c.GetInt("bookid")
	if err != nil {

		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("MainController.Get"), 307)
	}

	if err = user.(*models.User).LendBook(bookId); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	c.Redirect(c.URLFor("MainController.Get"), 307)
}

// ReturnBook ...
func (c *UserBookController) ReturnBook() {
	user := c.GetSession("user")

	flash := web.NewFlash()

	bookId, err := c.GetInt("bookid")
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("MainController.Get"), 307)
	}

	if err = user.(*models.User).ReturnBook(bookId); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}
	c.Redirect(c.URLFor("MainController.Get"), 307)
}
