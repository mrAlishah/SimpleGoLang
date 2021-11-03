package main

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)

	r.Run("localhost:8080")
}
