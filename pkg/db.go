package pkg

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func SetupDB() {
	db, err := sqlx.Connect("sqlite3", "./filebag.db")
	if err != nil {
		Log.Fatal(err)
	}

	DB = db

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			raw_url TEXT NOT NULL,
			file_name TEXT NOT NULL,
			file_size INTEGER NOT NULL,
			connection_count INTEGER NOT NULL,
			status TEXT NOT NULL,
			created_at INTEGER NOT NULL
		);
	`)

	if err != nil {
		Log.Fatal(err)
	}
}