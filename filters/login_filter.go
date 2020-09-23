package filters

import (
	"strings"

	"github.com/astaxie/beego/context"
)

// LoginFilter ...
var LoginFilter = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/index") ||
		strings.Compare(ctx.Input.URL(), "/") == 0 {
		return
	}

	user := ctx.Input.Session("user")

	if strings.HasPrefix(ctx.Input.URL(), "/login") && user != nil {
		ctx.Redirect(302, "/")
	}

	if user != nil ||
		strings.HasPrefix(ctx.Input.URL(), "/login") ||
		strings.HasPrefix(ctx.Input.URL(), "/book/reviews") {
		return
	}

	ctx.Redirect(302, "/")
}
