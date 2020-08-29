package controllers

import (
	"bookshelf/forms"

	"github.com/astaxie/beego"
)

// LoginController struct
type LoginController struct {
	beego.Controller
}

// Get Login function
func (c *LoginController) Get() {
	user := c.GetSession("user")
	if user != nil {
		c.Redirect("index", 307)
	}
	c.TplName = "login.tpl"
}

// Post Login function
func (c *LoginController) Post() {
	flash := beego.NewFlash()
	loginForm := forms.LoginForm{}
	if err := c.ParseForm(&loginForm); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("index", 303)
	}

	user, err := loginForm.GetData()
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}

	if err := user.DoLogin(); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		c.SetSession("user", user)
	}

	c.TplName = "login.tpl"
	c.Redirect("index", 303)
}

// Logout function
func (c *LoginController) Logout() {
	c.DestroySession()
	c.Redirect("index", 307)
}
