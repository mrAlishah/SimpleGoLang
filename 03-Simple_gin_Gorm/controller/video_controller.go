// Package controller implements api endpoint controllers
package controller

import (
	"log"
	"strconv"

	"simple_gin_gorm/entity"
	"simple_gin_gorm/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	Get(ctx *gin.Context) (entity.Video, error)
	FindAll() []entity.Video
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	srv service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &controller{
		srv: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.srv.FindAll()
}

func (c *controller) Get(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return video, err
	}
	video.ID = id
	return c.srv.Get(video), nil
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		log.Print("error in binding")
		return err
	}

	log.Print("video before save: ", video)
	c.srv.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	c.srv.Update(video)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id
	c.srv.Delete(video)
	return nil

}
