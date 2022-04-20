package controllers

import (
	"bookshelf/models"

	"github.com/beego/beego/v2/server/web"
)

// MainController struct
type MainController struct {
	web.Controller
}

// Get method
func (c *MainController) Get() {
	fd := web.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	user := c.GetSession("user")
	c.Data["IsUserLogged"] = user != nil
	if user != nil {
		u := user.(*models.User)
		if err := u.Read(); err != nil {
			c.DelSession("user")
			flash := web.NewFlash()
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect(c.URLFor("MainController.Get"), 307)
		}
		c.Data["UserId"] = u.Id
		c.Data["UserName"] = u.Name
		c.Data["UserRole"] = u.Role.Name
		c.Data["UserBooks"] = u.BooksLent
	}
	b := new(models.Book)
	c.Data["Books"] = b.All()
	c.TplName = "index.tpl"
}
