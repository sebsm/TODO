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

// div.footer {
//     position: fixed;
//     left: auto;
//     bottom: 0;
//     width: 100%;
//     color: inherit;
//     text-align: left;
//     background-color: #97a7b7;
//     font-size: 15px;
//     margin-bottom: 0px;
// }
// .menu {
//     margin-left: 10px;
//     margin-right: auto;
//     font-size: 18px;
// }
// .form-inline {
//     margin-top: 8px;
// }
// .navbar-inherit-top  {
//     position: relative;
//     min-height: 10px;
//     margin-top: 0px;
//     margin-bottom: 20px;
//     border: 1px solid transparent;
//     border-radius: 0px;
// }
// .html {
//     font-size: 18px;
//     -webkit-tap-highlight-color: rgba(0,0,0,0);
// }
// .marketing .col-lg-4 {
//     margin-bottom: 1.5rem;
//     text-align: center;
// }
// .marketing h2 {
//     font-weight: 400;
// }
// .marketing .col-lg-4 p {
//     margin-right: .75rem;
//     margin-left: .75rem;
// }
// table {
//     counter-reset: rowNumber ;
// }

// table tr td:first-child:before {
//     display: table-cell;
//     counter-increment: rowNumber;
//     content: counter(rowNumber) ".";
// }

//var input createtask
//completed := c.GetBool("completed")

//id, _ := strconv.Atoi(c.PostForm("id"))
//key := c.PostForm("id")
// if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 	return
// }
// request := &task{
// 	Title: title,
// }
// if err := db.Where("title = ?", title).First(&task).Error; err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 	return
// }\

// for i := 0; i < len(tasks); i++ {
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": fmt.Sprintf("%+v\n", tasks[i]),
// 	})
// }

// c.JSON(http.StatusOK, gin.H{
// 	"id":          db.Select("id").Where("title LIKE ?", s).Find(&tasks),
// 	"title":       db.Select("title").Where("title LIKE ?", s).Find(&tasks),
// 	"completed":   db.Select("completed").Where("title LIKE ?", s).Find(&tasks),
// 	"description": db.Select("description").Where("title LIKE ?", s).Find(&tasks),
// })

//fmt.Printf("ID: %v, Title: %v, CreatedAt: %v, Completed: %v, Description: %v \n", &tasks.ID, &tasks.Title, tasks.CreatedAt, tasks.Completed, tasks.Description)

// var name []string
// for i := 0; i < len(tasks); i++ {
// 	name[i] = fmt.Sprintf("%+v \n", tasks[i])
// }

//name := fmt.Sprintf("%+v \n", tasks)
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