package middleware

import (
	"github.com/astaxie/beego/context"
)

func FilterUser(ctx *context.Context) {
	_, ok := ctx.Input.Session("UserID").(int)

	if !ok && ctx.Request.RequestURI != "/loginAction" {

		ctx.Redirect(302, "/loginAction")
	}

}
