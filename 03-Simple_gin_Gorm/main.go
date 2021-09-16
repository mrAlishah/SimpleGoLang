package main

import (
	"log"
	"net/http"

	"simple_gin_gorm/controller"
	"simple_gin_gorm/repository"
	"simple_gin_gorm/service"

	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository("Test.db")
	videoService    service.VideoService       = service.NewVideoService(videoRepository)
	videoController controller.VideoController = controller.NewVideoController(videoService)
)

func main() {

	server := gin.Default()

	videos := server.Group("/videos")
	{
		videos.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		videos.GET("/:id", func(ctx *gin.Context) {
			video, err := videoController.Get(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"video": video})
			}
		})

		videos.POST("/", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}

		})

		videos.PUT("/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"video": "video updated successfully"})
			}
		})

		videos.DELETE("/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"video": "video deleted successfully"})
			}
		})
	}

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}

}
