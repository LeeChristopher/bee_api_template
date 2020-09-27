package middlewares

import (
	"bee_api_template/packages"
	"github.com/astaxie/beego/context"
	"strings"
)

/*
 * 认证中间件
 */
func Auth(ctx *context.Context) {
	response := &packages.CommonResponse{
		RunTime: 0,
		Code:    401,
		Message: "please login",
		Data:    make([]uint8, 0, 0),
	}
	authorization := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(strings.ToLower(authorization), "bearer ") {
		_ = ctx.Output.JSON(response, false, false)
		return
	}
	authTokenList := strings.Split(authorization, " ")
	if len(authTokenList) != 2 || authTokenList[0] != "Bearer" {
		_ = ctx.Output.JSON(response, false, false)
		return
	}

	_, err := packages.VerifyAuthToken(authTokenList[1])
	if err != nil {
		response.Message = "invalid token"
		_ = ctx.Output.JSON(response, false, false)
		return
	}
}
