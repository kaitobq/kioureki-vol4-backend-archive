package db

import (
	"backend/db/query"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var queryPaths = []string{
	"db/query/user.sql",
}

func ConnectDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_CONTAINER_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database successfully")

	time.Sleep(1 * time.Second)

	errors := query.ExecuteSQLFiles(db, queryPaths)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println("Error executing SQL file:", err)
		}
		log.Fatal("Failed to execute some SQL files")
	}

	fmt.Println("Executed SQL files successfully")

	return db
}