package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"SMS/src/internal/database"
)

type Student struct {
	Name       string `json:"name"`
	EnrollDate string `json:"enrollDate"`
}

func StudentRoutes(r *gin.Engine) {
	studentGroup := r.Group("/students")
	studentGroup.GET("/", getStudents)
	studentGroup.POST("/", createStudent)
}

func getStudents(c *gin.Context) {
	db := database.GetDB()

	var students []Student
	rows, err := db.Query("SELECT name, enrollDate FROM students")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting students",
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Name, &student.EnrollDate)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error getting student data",
			})
			return
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with student rows: %v", err)
		c.JSON(500, gin.H{"message": "Error processing student data"})
		return
	}

	c.JSON(200, students)
}

func createStudent(c *gin.Context) {
	db := database.GetDB()

	var student Student
	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating student BindJSON",
		})
		return
	}

	_, err = db.Exec("INSERT INTO students (name, enrollDate) VALUES (?, ?)", student.Name, student.EnrollDate)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error creating student",
		})
		return
	}

	c.JSON(200, student)
}
