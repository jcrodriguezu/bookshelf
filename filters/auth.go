package filters

import (
	"github.com/beego/beego/v2/server/web/context"
)

// AuthFilter ...
var AuthFilter = func(ctx *context.Context) {
	user := ctx.Input.Session("user")
	if user == nil {
		ctx.Redirect(302, "/")
	}
}
