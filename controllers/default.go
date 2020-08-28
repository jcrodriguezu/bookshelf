package controllers

import (
	"bookshelf/models"

	"github.com/astaxie/beego"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get method
func (c *MainController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["errors"] = flash.Data["error"]
	c.Data["notices"] = flash.Data["notice"]

	user := c.GetSession("user")
	c.Data["IsUserLogged"] = user != nil
	if user != nil {
		u := user.(*models.User)
		c.Data["UserName"] = u.Name
		c.Data["UserRole"] = u.Role.Name
		c.Data["UserBooks"] = u.BooksLent
	}
	b := new(models.Book)
	c.Data["Books"] = b.All()
	c.TplName = "index.tpl"
}
