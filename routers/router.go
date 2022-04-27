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

	bookNs := web.NewNamespace("/book",
		web.NSRouter("/get", &controllers.BookController{}),
		web.NSRouter("/new", &controllers.BookController{}, "post:New"),
		web.NSRouter("/edit", &controllers.BookController{}, "post:Edit"),
		web.NSRouter("/remove", &controllers.BookController{}, "get:Remove"),
		web.NSRouter("/lend", &controllers.UserBookController{}, "*:LendBook"),
		web.NSRouter("/return", &controllers.UserBookController{}, "*:ReturnBook"),
	)

	reviewNs := web.NewNamespace("/review",
		web.NSRouter("/get", &controllers.UserBookReviewController{}),
		web.NSRouter("/new", &controllers.UserBookReviewController{}, "post:New"),
		web.NSRouter("/list", &controllers.UserBookReviewController{}, "*:Reviews"),
	)

	web.Router("/scrap/:isbn", &controllers.BookController{}, "get:SearchIsbn")

	web.AddNamespace(bookNs, reviewNs)

}
