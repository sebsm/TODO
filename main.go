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

func findtasks(c *gin.Context) {
	var tasks []task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func routing(router *gin.Engine) {
	router.GET("/index", initial)
	router.GET("/tasks", findtasks)
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

	// dsn := "user=postgres password=s197328645S! dbname=todo port=5432 sslmode=disable TimeZone=Europe/Warsaw"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// Migrate the schema
	//db.Migrator().DropTable(&task{})
	//db.AutoMigrate(&task{})

	// Create
	//db.Create(&task{Description: "Wash the dishes", Completed: true})

	// Read
	// var product task
	// db.First(&product, 1)                        // find product with integer primary key
	// db.First(&product, "Description = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Completed", false)
	// // Update - update multiple fields
	// db.Model(&product).Updates(task{Completed: false, Description: "R2D2"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Completed": false, "Description": "R2D2"})

	// //Delete - delete product
	// db.Delete(&product, 1)
	connect()
	db.Create(&task{Description: "Wash the dishes", Completed: true})
	router := gin.Default()
	routing(router)

	router.Run(port)

	log.Fatal(autotls.Run(router))

}
