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
	fd := beego.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	user := c.GetSession("user")
	c.Data["IsUserLogged"] = user != nil
	if user != nil {
		u := user.(*models.User)
		if err := u.Read(); err != nil {
			c.DelSession("user")
			flash := beego.NewFlash()
			beego.Info(err)
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("index", 307)
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
