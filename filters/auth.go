package filters

import (
	"github.com/astaxie/beego/context"
)

// AuthFilter ...
var AuthFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil {
		ctx.Redirect(302, "/")
	}
}
