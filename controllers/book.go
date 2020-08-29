package controllers

import "github.com/astaxie/beego"

// BookController ...
type BookController struct {
	beego.Controller
}

// Get ...
func (c *BookController) Get() {
	c.TplName = "book.tpl"
}
