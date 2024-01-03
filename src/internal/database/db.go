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
	CreateStudentTables()
	CreateTeacherTables()
	CreateClassTables()
	CreateGradeTables()
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

func CreateStudentTables() {
	createStudentTable := `
		CREATE TABLE IF NOT EXISTS students(
			studentId INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			enrollDate VARCHAR(255) NOT NULL
		)
	`

	if _, err := DB.Exec(createStudentTable); err != nil {
		log.Fatalf("Failed to create students table: %s\n", err)
	}
}

func CreateTeacherTables() {
	createTeacherTable := `
		CREATE TABLE IF NOT EXISTS teachers(
			teacherId INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		)
	`

	if _, err := DB.Exec(createTeacherTable); err != nil {
		log.Fatalf("Failed to create teachers table: %s\n", err)
	}
}

func CreateClassTables() {
	createClassTable := `
		CREATE TABLE IF NOT EXISTS classes(
			classId INT AUTO_INCREMENT PRIMARY KEY,
			className VARCHAR(255) NOT NULL,
			classSection VARCHAR(255) NOT NULL,
			semester VARCHAR(255) NOT NULL,
			teacherId INT NOT NULL,
			studentsNumber INT NOT NULL,
			FOREIGN KEY (teacherId) REFERENCES teachers(teacherId)
		)
	`

	if _, err := DB.Exec(createClassTable); err != nil {
		log.Fatalf("Failed to create classes table: %s\n", err)
	}
}

func CreateGradeTables() {
	createGradeTables := `
		CREATE TABLE IF NOT EXISTS grades(
			studentId INT NOT NULL,
			classId INT NOT NULL,
			FOREIGN KEY (studentId) REFERENCES students(studentId),
			FOREIGN KEY (classId) REFERENCES classes(classId),
			grade DECIMAL(4,3) NOT NULL,
			PRIMARY KEY (studentId, classId)
		)
	`

	if _, err := DB.Exec(createGradeTables); err != nil {
		log.Fatalf("Failed to create grades table: %s\n", err)
	}
}

func GetDB() *sql.DB {
	return DB
}
