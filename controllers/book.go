package controllers

import (
	"bookshelf/forms"
	"bookshelf/models"

	"github.com/astaxie/beego"
)

// BookController ...
type BookController struct {
	beego.Controller
}

// Get ...
func (c *BookController) Get() {
	fd := beego.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	user := c.GetSession("user")
	if user == nil {
		c.Redirect("index", 307)
	}

	c.Data["Form"] = &forms.BookForm{}
	c.TplName = "book.tpl"
}

// Post new book
func (c *BookController) Post() {
	flash := beego.NewFlash()

	bookForm := forms.BookForm{}
	if err := c.ParseForm(&bookForm); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("index", 303)
	}

	book, err := bookForm.GetData()
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := book.Insert(); err != nil {
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			flash.Notice("Book successful created")
			flash.Store(&c.Controller)
		}
	}
	c.Redirect("book", 303)
}

// Remove ...
func (c *BookController) Remove() {
	user := c.GetSession("user")
	if user != nil {
		c.Redirect("/index", 307)
	}

	flash := beego.NewFlash()

	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		book := &models.Book{Id: id}
		if err := book.Delete(); err != nil {
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		}
	}
	c.Redirect("/index", 303)
}
