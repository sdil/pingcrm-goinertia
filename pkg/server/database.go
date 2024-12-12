package server

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)
  
func ConnectDb() *sql.DB {
	db, err := sql.Open("sqlite3", "pingcrm.db?_journal_mode=wal")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}