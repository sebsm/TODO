package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

// const (
// 	port = "8080"
// )

func initial(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home page",
	})
}

// func invalid(c *gin.Context) {
// 	c.JSON(404, gin.H{
// 		"message": "Unable to proceed",
// 	})
// }
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
	router.GET("/", initial)
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
	// router.GET("/findtask", func(c *gin.Context) {
	// 	c.HTML(200, "find.html", gin.H{
	// 		"title": "Find task",
	// 	})
	// })

	//router.NoRoute(invalid)
}

var db *gorm.DB

type Serwer struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func connect() {

	//dsn := "host=ec2-54-196-65-186.compute-1.amazonaws.com user=wmvufhsdqlrtkr password=6e0f3c5109cb6e4a42e02e6924d2e3ee0f7eb36774a379df9b0072871dfb91f2 dbname=d5pffg8tbeebjo port=5432 sslmode=disable"
	//dst := "postgres://wmvufhsdqlrtkr:6e0f3c5109cb6e4a42e02e6924d2e3ee0f7eb36774a379df9b0072871dfb91f2@ec2-54-196-65-186.compute-1.amazonaws.com:5432/d5pffg8tbeebjo"
	// dsn := "user= password= dbname= port= sslmode= TimeZone=Europe/Warsaw"
	//database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	//dsn := os.Getenv("DATABASE_URL")

	// sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	// if err == nil {
	// 	fmt.Sprintln(err)
	// }
	// database, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&task{})

	db = database

}

func main() {

	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")
	log.Print(port)
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	log.Print(os.Getenv("DATABASE_URL"))
	connect()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Static("/css", "../assets/css")
	router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	routing(router)
	http.ListenAndServe(port, router)
	//log.Fatal(autotls.Run(router))
	//router.Run(":" + port)
	//router.Run(port)

}
