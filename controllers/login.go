package controllers

import (
	"bookshelf/forms"
	"bookshelf/models"

	"github.com/beego/beego/v2/server/web"
)

// LoginController struct
type LoginController struct {
	web.Controller
}

// Get Login function
func (c *LoginController) Get() {
	fd := web.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	c.Data["Form"] = &forms.LoginForm{}
	c.TplName = "login.tpl"
}

// Post Login function
func (c *LoginController) Post() {
	flash := web.NewFlash()

	loginForm := &forms.LoginForm{}
	if err := c.ParseForm(loginForm); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	}

	user, err := forms.ToModel(loginForm)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := user.(*models.User).DoLogin(); err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			c.SetSession("user", user)
			c.Redirect(c.URLFor("MainController.Get"), 303)
			return
		}
	}

	c.Redirect(c.URLFor("LoginController.Get"), 303)
}

// Logout function
func (c *LoginController) Logout() {
	c.DestroySession()
	c.Redirect(c.URLFor("MainController.Get"), 307)
}
