package controllers

import (
	"bookshelf/forms"
	"bookshelf/models"

	"github.com/astaxie/beego"
)

// LoginController struct
type LoginController struct {
	beego.Controller
}

// Get Login function
func (c *LoginController) Get() {
	fd := beego.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	c.Data["Form"] = &forms.LoginForm{}
	c.TplName = "login.tpl"
}

// Post Login function
func (c *LoginController) Post() {
	flash := beego.NewFlash()

	loginForm := &forms.LoginForm{}
	if err := c.ParseForm(loginForm); err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}

	user, err := forms.ToModel(loginForm)
	if err != nil {
		beego.Info(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := user.(*models.User).DoLogin(); err != nil {
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			c.SetSession("user", user)
			c.Redirect("index", 303)
		}
	}

	c.Redirect("login", 303)
}

// Logout function
func (c *LoginController) Logout() {
	c.DestroySession()
	c.Redirect("index", 307)
}
