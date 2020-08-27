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
