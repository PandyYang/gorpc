package controller

import (
	"github.com/kataras/iris/v12/mvc"
	"gorpc/iris/repository"
	"gorpc/iris/service"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepo := repository.NewMovieManager()
	movie := service.NewMovieServiceManager(movieRepo)
	res := movie.ShowMovieName()

	return mvc.View{
		Name: "movie/index.html",
		Data: res,
	}
}
