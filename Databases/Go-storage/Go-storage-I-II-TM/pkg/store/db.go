package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Init() *sql.DB {
	dataSource := "root:@tcp(localhost:3306)/storage"

	StorageDB, err := sql.Open("mysql", dataSource)

	if err != nil {
		fmt.Println("ERROR 1")
	}

	if err = StorageDB.Ping(); err != nil {
		fmt.Println("ERROR DE PING")
	}

	log.Println("database configured")

	return StorageDB
}

// Otra forma de inicializar la conexion

func connectToDatabse() (engine *gin.Engine, db *sql.DB) {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	db, err = sql.Open("mysql", configDB.FormatDSN())

	if err != nil {
		panic(err)
	}

	return engine, db
}
