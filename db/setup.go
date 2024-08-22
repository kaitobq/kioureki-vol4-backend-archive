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
	"golang.org/x/crypto/bcrypt"
)

var queryPaths = []string{
	"db/query/user.sql",
	"db/query/organization.sql",
	"db/query/user_organization_invitation.sql",
	"db/query/user_organization_membership.sql",
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

	// insertMockData(db)

	fmt.Println("Executed SQL files successfully")

	return db
}

func insertMockData(db *sql.DB) {
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("user%d", i)
		email := fmt.Sprintf("user%d@mail.com", i)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, string(hashedPassword))
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("organization%d", i)
		_, err := db.Exec("INSERT INTO organizations (name) VALUES (?)", name)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 1; i <= 10; i++ {
		_, err := db.Exec("INSERT INTO user_organization_memberships (organization_id, user_id) VALUES (?, ?)", i, i)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Inserted mock data successfully")
}