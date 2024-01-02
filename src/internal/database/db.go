package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	ConnectDB()
	CreateTables()
}

func ConnectDB() {
	var err error

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbhost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, dbhost, dbName)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %s\n", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting database: %s\n", err)
	}

	log.Println("Connected to Database")
}

func CreateTables() {
	createStudentTable := `
		CREATE TABLE IF NOT EXISTS students(
			studentId INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			enrollDate VARCHAR(255)
		)
	`

	if _, err := DB.Exec(createStudentTable); err != nil {
		log.Fatalf("Failed to create students table: %s\n", err)
	}
}

func GetDB() *sql.DB {
	return DB
}
