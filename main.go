package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type task struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
}

type createtask struct {
	Title       string    `json:"title" binding:"title"`
	CreatedAt   time.Time `json:"createdat" binding:"createdat"`
	UpdatedAt   time.Time `json:"updatedat" binding:"updatedat"`
	Completed   bool      `json:"completed" binding:"completed"`
	Description string    `json:"description" binding:"description"`
}

const (
	port = ":8080"
)

func initial(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi, friend",
	})
	return
}

func invalid(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Unable to proceed",
	})
	return
}
func addtasks(c *gin.Context) {
	var example createtask

	if err := c.ShouldBindJSON(&example); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := task{
		Title:       example.Title,
		CreatedAt:   example.UpdatedAt,
		UpdatedAt:   example.UpdatedAt,
		Completed:   example.Completed,
		Description: example.Description,
	}

	db.Create(&request)

	c.JSON(http.StatusOK, gin.H{"data": request})

}

func findtasks(c *gin.Context) {
	var tasks []task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func routing(router *gin.Engine) {
	router.GET("/home", initial)
	router.GET("/tasks", findtasks)
	router.POST("/add", addtasks)
	router.NoRoute(invalid)
}

var db *gorm.DB

func connect() {
	dsn := "user=postgres password=s197328645S! dbname=todo port=5432 sslmode=disable TimeZone=Europe/Warsaw"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&task{})

	db = database
}

func main() {
	db.Migrator().DropTable(&task{})
	connect()
	db.Create(&task{Description: "Wash the dishes", Completed: true})
	router := gin.Default()
	routing(router)

	router.Run(port)

	log.Fatal(autotls.Run(router))

}
