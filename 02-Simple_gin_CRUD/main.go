package main

import (
	"log"
	"net/http"
	"simple_gin_crud/controller"
	"simple_gin_crud/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default() //gin router ~ Server

	videoRoutes := server.Group("/video")
	{
		videoRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		videoRoutes.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.Save(ctx))
		})
	}

	//server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := server.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
