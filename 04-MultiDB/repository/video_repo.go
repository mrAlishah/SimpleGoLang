package repository

import (
	"errors"
	"multiDB/entity"
)

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

type VideoRepository interface {
	GetAll() ([]entity.Video, error)
	GetVideoByTitle(string) (entity.Video, error)
	Add(entity.Video) error
	Update(entity.Video, int) error
}

var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

//factory function
func GetDatabaseHandler(dbtype uint8, connection string) (VideoRepository, error) {

	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	case POSTGRESQL:
		return NewPQHandler(connection)
	}
	return nil, DBTypeNotSupported
}
