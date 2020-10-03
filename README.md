# TODO

# Simple TODO app using Postgres + GORM + Gin framework

1. Go to [Go website](https://golang.org/) and set up correct paths.

2. Configure PostgreSQL (or other DB, but you must then change database driver) and pass it to these lines.

```
dsn := "user=your_user(example user:ard) password=your_pass(your password example password=1234) dbname=your_name (example dbname:abc) port=your_port (example: port=1234) sslmode=disable TimeZone=Europe/Warsaw"
database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

3. Connect to your database

4. Enter the main.go directory via cmd (Windows) and enter command:

```
go run main.go
```

5. Enter http://localhost:8080/home in internet browser and test 


# Important links

1. [Gin framework](https://github.com/gin-gonic/gin)
2. [GORM ORM](https://github.com/go-gorm/gorm)
3. Icons:
  * Add https://thenounproject.com/term/add-task/945026/
  * Delete https://thenounproject.com/search/?q=DELETE+TASK&i=525770
  * Find https://thenounproject.com/search/?q=FIND+TASK&i=532103
  * Update https://thenounproject.com/search/?q=UPDATE+TASK&i=2676042
  * All https://thenounproject.com/term/task/3129877/
4. PostgreSQL https://www.postgresql.org/

# Images

![Main page](https://raw.githubusercontent.com/selvert/TODO/master/images/Main.PNG)

![Add page](https://raw.githubusercontent.com/selvert/TODO/master/images/Add.PNG)

![Delete page](https://raw.githubusercontent.com/selvert/TODO/master/images/Delete.PNG)

![Find page](https://raw.githubusercontent.com/selvert/TODO/master/images/Find.PNG)

![Update page](https://raw.githubusercontent.com/selvert/TODO/master/images/Update.PNG)

![All page](https://raw.githubusercontent.com/selvert/TODO/master/images/All.PNG)
