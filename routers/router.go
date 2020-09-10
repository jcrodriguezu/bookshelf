package routers

import (
	"bookshelf/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
	beego.Router("/book", &controllers.BookController{})
	beego.Router("/book/new", &controllers.BookController{}, "post:New")
	beego.Router("/book/edit", &controllers.BookController{}, "post:Edit")
	beego.Router("/book/remove", &controllers.BookController{}, "get:Remove")
	beego.Router("/user/book/lend", &controllers.UserBookController{}, "*:LendBook")
	beego.Router("/user/book/return", &controllers.UserBookController{}, "*:ReturnBook")
}
