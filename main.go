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
	Title string `json:"title" binding:"title"`
	// CreatedAt   time.Time `json:"createdat" binding:"createdat"`
	// UpdatedAt   time.Time `json:"updatedat" binding:"updatedat"`
	Completed   bool   `json:"completed" binding:"completed"`
	Description string `json:"description" binding:"description"`
}

type changetask struct {
	Title string `json:"title" `
	// CreatedAt   time.Time `json:"createdat"`
	// UpdatedAt   time.Time `json:"updatedat"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
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
	var input createtask

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := task{
		Title: input.Title,
		// CreatedAt:   input.CreatedAt,
		// UpdatedAt:   input.UpdatedAt,
		Completed:   input.Completed,
		Description: input.Description,
	}

	db.Create(&request)

	c.JSON(http.StatusOK, gin.H{"data": request})

}

func findtask(c *gin.Context) {
	var task task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func findtasks(c *gin.Context) {
	var tasks []task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func updatetask(c *gin.Context) {
	var task task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input changetask
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func deletetask(c *gin.Context) {
	// Get model if exist
	var task task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func routing(router *gin.Engine) {
	router.GET("/home", initial)
	router.GET("/tasks", findtasks)
	router.POST("/add", addtasks)
	router.GET("/tasks/:id", findtask)
	router.PATCH("/tasks/:id", updatetask)
	router.DELETE("/books/:id", deletetask)
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

	connect()
	//db.Create(&task{Description: "Wash the dishes", Completed: true})
	router := gin.Default()
	routing(router)
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	//router.Run(port)
	http.ListenAndServe(port, router)
	log.Fatal(autotls.Run(router))

}
