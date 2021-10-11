package repository

import (
	"multiDB/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

func (dbHandler *MongodbHandler) GetAll() ([]entity.Video, error) {
	s := dbHandler.getFreshSession()
	defer s.Close()
	videos := []entity.Video{}
	err := s.DB("Test").C("videos").Find(nil).All(&videos)
	return videos, err
}

func (dbHandler *MongodbHandler) GetVideoByTitle(title string) (entity.Video, error) {
	s := dbHandler.getFreshSession()
	defer s.Close()
	r := entity.Video{}
	err := s.DB("Test").C("videos").Find(bson.M{"title": title}).One(&r)
	return r, err
}

func (dbHandler *MongodbHandler) Add(video entity.Video) error {
	s := dbHandler.getFreshSession()
	defer s.Close()
	return s.DB("Test").C("videos").Insert(video)
}

func (dbHandler *MongodbHandler) Update(video entity.Video, id int) error {
	s := dbHandler.getFreshSession()
	defer s.Close()
	return s.DB("Test").C("videos").Update(bson.M{"id": id}, video)
}

func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
