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

	bookform := &forms.BookForm{}
	c.Data["Action"] = "BookController.New"

	id, err := c.GetInt("id")
	if err == nil {
		book := &models.Book{Id: id}
		if err := book.GetById(id); err != nil {
			flash := beego.NewFlash()
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			c.Data["Action"] = "BookController.Edit"
			bookform = &forms.BookForm{
				Id:     book.Id,
				Title:  book.Title,
				Author: book.Author,
				Copies: book.Copies,
			}
		}
	}

	c.Data["Form"] = bookform
	c.TplName = "book.tpl"
}

// New book
func (c *BookController) New() {
	flash := beego.NewFlash()

	bookForm := forms.BookForm{}
	if err := c.ParseForm(&bookForm); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/index", 303)
	}

	book, err := bookForm.ToModel()
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
	c.Redirect("/book", 303)
}

// Edit book
func (c *BookController) Edit() {
	user := c.GetSession("user")
	if user == nil {
		c.Redirect("index", 307)
	}

	flash := beego.NewFlash()

	bookForm := forms.BookForm{}
	if err := c.ParseForm(&bookForm); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("index", 303)
	}

	book, err := bookForm.ToModel()
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := book.Update(); err != nil {
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			flash.Notice("Book successful updated")
			flash.Store(&c.Controller)
		}
	}
	c.Redirect("/index", 303)
}

// Remove book
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
