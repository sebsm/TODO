package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type task struct {
	ID          uint      `form:"id" json:"id" gorm:"primary_key"`
	Title       string    `form:"title" json:"title"`
	CreatedAt   time.Time `form:"createdat" json:"createdat"`
	UpdatedAt   time.Time `form:"updatedat" json:"updatedat"`
	Completed   bool      `form:"completed" json:"completed"`
	Description string    `form:"description" json:"description"`
}

// type login struct {
// 	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
// 	Password string `form:"password" json:"password" xml:"password" binding:"required"`
// }

const (
	PORT = ":8080"
)

func initial(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"title": "Home page",
	})
}

func invalid(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Unable to proceed",
	})
}
func addtasks(c *gin.Context) {

	title := c.PostForm("title")
	completed, _ := strconv.Atoi(c.PostForm("completed"))

	description := c.PostForm("description")
	var finished bool
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title": "Add task",
	})
	if completed == 1 {
		finished = true
	} else {
		finished = false
	}
	request := &task{
		Title:       title,
		Completed:   finished,
		Description: description,
	}

	db.Create(&request)

	c.JSON(http.StatusOK, gin.H{"data": request})

}

func deletetask(c *gin.Context) {
	title := c.PostForm("title")

	c.HTML(http.StatusOK, "delete.html", gin.H{
		"title": "Delete task",
	})

	request := &task{
		Title: title,
	}

	db.Where("title = ?", title).Delete(&task{})

	c.JSON(http.StatusOK, gin.H{"data": request})
}

func updatetask(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	name := c.PostForm("name")
	var finished = false

	c.HTML(http.StatusOK, "update.html", gin.H{
		"title": "Update task",
	})

	switch completed {
	case 0:
		finished = false
	case 1:
		finished = true
	}
	request := &task{
		Title:       title,
		Completed:   finished,
		Description: description,
	}
	db.Model(&task{}).Where("title = ?", name).Updates(map[string]interface{}{"Title": request.Title, "Completed": request.Completed, "Description": request.Description})
	c.JSON(http.StatusOK, gin.H{"data": request})
}
func findtask(c *gin.Context) {
	var tasks []task
	title := c.PostForm("title")
	merged := "%" + title + "%"
	db.Where("title LIKE ?", merged).Find(&tasks)

	c.HTML(200, "find.html", gin.H{
		"title": "Find task",
		"name":  fmt.Sprintf("%+v\n", tasks),
		"tasks": tasks,
	})
	text := fmt.Sprintln(merged)
	io.WriteString(os.Stdout, text)

}

func findtasks(c *gin.Context) {
	var tasks []task
	db.Find(&tasks)
	c.HTML(200, "tasks.html", gin.H{
		"title": "All tasks",
		"tasks": tasks,
	})

	s := fmt.Sprintln(tasks)
	io.WriteString(os.Stdout, s)
}

func routing(router *gin.Engine) {
	router.GET("/home", initial)
	router.GET("/tasks", findtasks)
	router.POST("/findtask", findtask)
	router.POST("/addtask", addtasks)
	router.POST("/deletetask", deletetask)
	router.POST("/updatetask", updatetask)

	router.GET("/addtask", func(c *gin.Context) {
		c.HTML(200, "add.html", gin.H{
			"title": "Add task",
		})
	})
	router.GET("/deletetask", func(c *gin.Context) {
		c.HTML(200, "delete.html", gin.H{
			"title": "Delete task",
		})
	})
	router.GET("/updatetask", func(c *gin.Context) {
		c.HTML(200, "update.html", gin.H{
			"title": "Update task",
		})
	})
	router.GET("/findtask", func(c *gin.Context) {
		c.HTML(200, "find.html", gin.H{
			"title": "Find task",
		})
	})

	router.NoRoute(invalid)
}

var db *gorm.DB

func connect() {
	// dsn := "user= password= dbname= port= sslmode= TimeZone=Europe/Warsaw"
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&task{})

	db = database

}

func main() {

	connect()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Static("/css", "../assets/css")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routing(router)
	http.ListenAndServe(PORT, router)
	log.Fatal(autotls.Run(router))

}
