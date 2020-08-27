package controllers

import (
	"bookshelf/models"
	"fmt"

	"github.com/astaxie/beego"
)

// LoginController struct
type LoginController struct {
	beego.Controller
}

type loginForm struct {
	Username interface{} `form:"username"`
	Password interface{} `form:"password"`
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
	u := loginForm{}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("Error parsing form user login")
		c.Redirect("index", 303)
	}
	fmt.Println(u)

	user, err := new(models.User).GetByUserNameAndPassword(u.Username.(string), u.Password.(string))
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
		c.Redirect("index", 303)
	} else {
		c.SetSession("user", user)
		c.Redirect("index", 303)
	}
	c.TplName = "login.tpl"
}

// Logout function
func (c *LoginController) Logout() {
	c.DestroySession()
	c.Redirect("index", 307)
}
