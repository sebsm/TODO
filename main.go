package main

import (
	"log"
	"net/http"
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

type createtask struct {
	Title string `form:"title" binding:"required"`
	// CreatedAt   time.Time `json:"createdat" binding:"createdat"`
	// UpdatedAt   time.Time `json:"updatedat" binding:"updatedat"`
	Completed   bool   `form:"completed" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type changetask struct {
	Title string `json:"title" `
	// CreatedAt   time.Time `json:"createdat"`
	// UpdatedAt   time.Time `json:"updatedat"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

const (
	port = ":8081"
)

func initial(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"title": "Home page",
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
	//var input createtask
	title := c.PostForm("title")
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	//completed := c.GetBool("completed")
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

	// if err := c.ShouldBind(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	request := &task{
		Title:       title,
		Completed:   finished,
		Description: description,
	}

	db.Create(&request)

	c.JSON(http.StatusOK, gin.H{"data": request})

}

func deletetask(c *gin.Context) {
	// Get model if exist
	title := c.PostForm("title")
	//id, _ := strconv.Atoi(c.PostForm("id"))
	//key := c.PostForm("id")
	// if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }
	c.HTML(http.StatusOK, "delete.html", gin.H{
		"title": "Delete task",
	})

	//value, _ := strconv.ParseUint(key, 8, 0)
	request := &task{
		//ID:    value,
		Title: title,
	}

	db.Where("title = ?", title).Delete(&task{})

	c.JSON(http.StatusOK, gin.H{"data": request})
}

func updatetask(c *gin.Context) {
	//var task task
	// Validate input
	//var input changetask
	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }
	//var request task
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

	// if completed == 1 {
	// 	finished = true
	// } else if completed == 0 {
	// 	finished = false
	// } else {
	// 	finished = false
	// }

	request := &task{
		Title:       title,
		Completed:   finished,
		Description: description,
	}

	//db.Model(&task{}).Where("title = ?", name).Updates(task{Title: request.Title, Completed: request.Completed, Description: request.Description})
	db.Model(&task{}).Where("title = ?", name).Updates(map[string]interface{}{"Title": request.Title, "Completed": request.Completed, "Description": request.Description})
	c.JSON(http.StatusOK, gin.H{"data": request})
}
func findtask(c *gin.Context) {
	title := c.PostForm("title")
	c.HTML(200, "find.html", gin.H{
		"title": "Find task",
	})
	request := &task{
		Title: title,
	}
	// if err := db.Where("title = ?", title).First(&task).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }\
	db.Where("title = ?", title).First(&request)
	c.JSON(http.StatusOK, gin.H{"data": request})

}

func findtasks(c *gin.Context) {
	var tasks []task
	db.Find(&tasks)
	c.HTML(200, "tasks.html", gin.H{
		"title": "All tasks",
	})
	c.JSON(http.StatusOK, gin.H{"data": tasks})
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
	router.LoadHTMLGlob("templates/*")
	routing(router)
	// s := &http.Server{
	// 	Addr:           ":8081",
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
