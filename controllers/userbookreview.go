package controllers

import (
	"bookshelf/forms"
	"bookshelf/models"

	"github.com/astaxie/beego"
)

// UserBookReviewController ...
type UserBookReviewController struct {
	beego.Controller
}

// Reviews ...
func (c *UserBookReviewController) Reviews() {
	flash := beego.NewFlash()
	bookId, err := c.GetInt("bookid")
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/index", 307)
	}

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
		c.Redirect("index", 307)
	}

	fd := beego.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	reviewform := &forms.ReviewForm{}
	c.Data["Action"] = "UserBookReviewController.New"

	bookid, err := c.GetInt("bookid")
	if err != nil {
		flash := beego.NewFlash()
		beego.Info(err)
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
	flash := beego.NewFlash()

	reviewForm := &forms.ReviewForm{}
	if err := c.ParseForm(reviewForm); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/index", 303)
	}

	review, err := forms.ToModel(reviewForm)
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := review.(*models.Review).Insert(); err != nil {
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			flash.Notice("Review successful created")
			flash.Store(&c.Controller)
		}
	}

	c.Redirect("/index", 303)
}
