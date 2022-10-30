package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorpc/iris/controller"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./iris/views", ".html"))

	// 注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controller.MovieController))
	app.Run(iris.Addr("localhost:8080"))
}
