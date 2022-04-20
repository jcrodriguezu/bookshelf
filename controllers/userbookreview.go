package controllers

import (
	"bookshelf/forms"
	"bookshelf/models"

	"github.com/beego/beego/v2/server/web"
)

// UserBookReviewController ...
type UserBookReviewController struct {
	web.Controller
}

// Reviews ...
func (c *UserBookReviewController) Reviews() {
	flash := web.NewFlash()
	bookId, err := c.GetInt("bookid")
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("MainController.Get"), 307)
	}

	user := c.GetSession("user")
	c.Data["IsUserLogged"] = user != nil

	rv := &models.Review{}
	reviews := rv.AllByBook(bookId)
	c.Data["Id"] = bookId
	c.Data["Reviews"] = reviews
	c.TplName = "reviews.tpl"
}

// Get ...
func (c *UserBookReviewController) Get() {
	user := c.GetSession("user")
	if user == nil {
		c.Redirect(c.URLFor("MainController.Get"), 307)
	}

	fd := web.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	reviewform := &forms.ReviewForm{}
	c.Data["Action"] = "UserBookReviewController.New"

	bookid, err := c.GetInt("bookid")
	if err != nil {
		flash := web.NewFlash()
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}

	reviewform.BookId = bookid
	reviewform.UserId = user.(*models.User).Id

	c.Data["Form"] = reviewform
	c.TplName = "form.tpl"
}

// New ...
func (c *UserBookReviewController) New() {
	flash := web.NewFlash()

	reviewForm := &forms.ReviewForm{}
	if err := c.ParseForm(reviewForm); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("MainController.Get"), 303)
	}

	review, err := forms.ToModel(reviewForm)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := review.(*models.Review).Insert(); err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			flash.Notice("Review successful created")
			flash.Store(&c.Controller)
		}
	}

	c.Redirect(c.URLFor("MainController.Get"), 303)
}
