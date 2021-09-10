package main

import (
	"log"
	"net/http"
	"simple_gin_server/controllers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User

func main() {
	r := gin.Default() //gin router ~ Server

	//Simple get route handler
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	//Simple get handler by function . attention to func getHandeler without ()
	r.GET("/msg", getHandler)

	//send string
	r.GET("/msg1", getHandler1)

	//use controler func
	r.GET("/msg2", controllers.Get)

	//simple crud
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUsers)
		userRoutes.POST("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

//------------------------------------------
func getHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":  true,
		"message2": "hello world!",
	})
}

func getHandler1(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hi Mostafa")
}

//-----------------  CRUD  --------------------
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
}

func CreateUsers(c *gin.Context) {
	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body.",
		})
		return
	}

	reqBody.ID = uuid.New().String()

	Users = append(Users, reqBody)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
	})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")

	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body.",
		})
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(http.StatusOK, gin.H{
				"error": false,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			//t := [1, 431, 44, 546, 61]
			//t index [0, 1, 2, 3, 4]
			//t[:2] == [1, 431]
			//t[2 + 1:] == [546, 61]
			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"error": false,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})

}

//------------------------------------------
