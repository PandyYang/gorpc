package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorpc/iris/backend/web/controller"
)

func main() {

	// 创建iris实例
	app := iris.New()
	// 设置错误模式
	app.Logger().SetLevel("debug")
	// 注册模板
	app.RegisterView(iris.HTML("./iris/backend/web/views", ".html").Layout(
		"shared/layout.html").Reload(true))

	app.HandleDir("/assets", "./iris/backend/web/assets")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// 注册控制器 实现路由
	mvc.New(app.Party("/hello")).Handle(new(controller.MovieController))

	// 启动服务
	app.Run(iris.Addr("localhost:8080"))
}
