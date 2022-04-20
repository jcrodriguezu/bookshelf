package routers

import (
	"bookshelf/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/index", &controllers.MainController{})
	web.Router("/login", &controllers.LoginController{})
	web.Router("/logout", &controllers.LoginController{}, "*:Logout")
	web.Router("/book/get", &controllers.BookController{})
	web.Router("/book/new", &controllers.BookController{}, "post:New")
	web.Router("/book/edit", &controllers.BookController{}, "post:Edit")
	web.Router("/book/remove", &controllers.BookController{}, "get:Remove")
	web.Router("/book/lend", &controllers.UserBookController{}, "*:LendBook")
	web.Router("/book/return", &controllers.UserBookController{}, "*:ReturnBook")
	web.Router("/review/get", &controllers.UserBookReviewController{})
	web.Router("/review/new", &controllers.UserBookReviewController{}, "post:New")
	web.Router("/review/list", &controllers.UserBookReviewController{}, "*:Reviews")
}
