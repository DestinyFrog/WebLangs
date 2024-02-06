package db

import (
	"database/sql"
	
	"app/config"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", config.Config.Database.Url )
}