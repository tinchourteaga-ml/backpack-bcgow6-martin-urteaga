package store

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
