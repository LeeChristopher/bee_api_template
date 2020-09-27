package routers

import (
	"bee_api_template/controllers"
	"bee_api_template/middlewares"
	"github.com/astaxie/beego"
)

func init() {
	namespace := beego.NewNamespace("/api",
		//登陆
		beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
		//用户模块
		beego.NSNamespace("/users",
			beego.NSBefore(middlewares.Auth),
			beego.NSRouter("/", &controllers.UserController{}, "get:Index"),
			beego.NSRouter("/", &controllers.UserController{}, "post:Add"),
			beego.NSRouter("/:id([1-9][0-9]*)", &controllers.UserController{}, "get:Info"),
			beego.NSRouter("/:id([1-9][0-9]*)", &controllers.UserController{}, "put:Update"),
			beego.NSRouter("/:id([1-9][0-9]*)", &controllers.UserController{}, "delete:Delete"),
		),
	)
	beego.AddNamespace(namespace)
}
