package controllers

import (
	"bookshelf/forms"

	"github.com/astaxie/beego"
)

// ReviewController ...
type ReviewController struct {
	beego.Controller
}

// Get ...
func (c *ReviewController) Get() {
	fd := beego.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	reviewform := &forms.ReviewForm{}
	c.Data["Action"] = "ReviewController.New"

	c.Data["Form"] = reviewform
	c.TplName = "form.tpl"
}
