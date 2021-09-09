package main

import (
	"log"
	"net/http"
	"simple_gin_server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //gin router ~ Server

	r.GET("/", getHandler)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping1", controllers.Get)

	r.GET("/ping2", getHandler1)

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func getHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  true,
		"message2": "hello world!",
	})
}

func getHandler1(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hi Mostafa")
}
