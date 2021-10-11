package repository

import (
	"database/sql"
	"fmt"
	"log"
	"multiDB/entity"
)

type SQLHandler struct {
	*sql.DB
}

func (dbHandler *SQLHandler) GetAll() ([]entity.Video, error) {
	return dbHandler.sendQuery("select * from Videos")
}

func (dbHandler *SQLHandler) GetVideoByTitle(title string) (entity.Video, error) {

	row := dbHandler.QueryRow(fmt.Sprintf("select * from videos where title = '%s'", title)) //? for mysql or sqlite and it used to be $1 for pq
	r := entity.Video{}
	err := row.Scan(&r.ID, &r.Title, &r.Description, &r.Url)
	return r, err
}

func (dbHandler *SQLHandler) Add(video entity.Video) error {
	_, err := dbHandler.Exec(fmt.Sprintf("Insert into videos (title,description,url) values ('%s','%s','%s')", video.Title, video.Description, video.Url))
	return err
}

func (dbHandler *SQLHandler) Update(video entity.Video, id int) error {
	_, err := dbHandler.Exec(fmt.Sprintf("Update videos set title = '%s' ,description = '%s',url = '%s' where id = %d", video.Title, video.Description, video.Url, id))
	return err
}

func (dbHandler *SQLHandler) sendQuery(q string) ([]entity.Video, error) {
	Videos := []entity.Video{}
	rows, err := dbHandler.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := entity.Video{}
		err := rows.Scan(&r.ID, &r.Title, &r.Description, &r.Url)
		if err != nil {
			log.Println(err)
			continue
		}
		Videos = append(Videos, r)
	}

	return Videos, rows.Err()
}
