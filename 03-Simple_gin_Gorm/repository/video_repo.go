package repository

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"simple_gin_gorm/entity"
)

type VideoRepository interface {
	FindAll() []entity.Video
	Get(video entity.Video) entity.Video

	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository(fileName string) VideoRepository {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(&entity.Video{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	return &database{
		connection: db,
	}
}

func (db *database) Save(video entity.Video) {
	log.Print("video before repo save: ", video)
	result := db.connection.Clauses(clause.OnConflict{DoNothing: true}).Create(&video)
	if result.Error != nil {
		log.Print("db save error: ", result.Error.Error())
	}
}

func (db *database) Get(video entity.Video) entity.Video {
	result := db.connection.First(&video)
	if result.Error != nil {
		log.Print("db get error: ", result.Error.Error())
	}
	return video
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	result := db.connection.Find(&videos)
	if result.Error != nil {
		log.Print("db find all error: ", result.Error.Error())
	}
	return videos
}

func (db *database) Update(video entity.Video) {
	result := db.connection.Save(&video)
	if result.Error != nil {
		log.Print("db update error: ", result.Error.Error())
	}
}

func (db *database) Delete(video entity.Video) {
	result := db.connection.Delete(&video)
	if result.Error != nil {
		log.Print("db delete error: ", result.Error.Error())
	}
}
