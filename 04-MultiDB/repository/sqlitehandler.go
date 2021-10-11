package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteHandler struct {
	*SQLHandler
}

func NewSQLiteHandler(connection string) (*SQLiteHandler, error) {
	db, err := sql.Open("sqlite3", connection)
	// createdatabase(db)
	return &SQLiteHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}

func createdatabase(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS
						videos(id INTEGER PRIMARY KEY AUTOINCREMENT, 
                        title TEXT,
                        description TEXT,
                        url TEXT) `)

}
