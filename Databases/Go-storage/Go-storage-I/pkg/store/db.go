package store

import (
	"database/sql"
	"log"
)

var StorageDB *sql.DB

func init() {
	dataSource := "root:@tcp(localhost.3306)/storage"

	StorageDB, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("database configured")
}
