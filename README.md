# TODO

# Simple TODO app using Postgres + GORM + Gin framework

1. Go to [Go website](https://golang.org/) and set up correct paths.

2. Configure PostgreSQL (or other DB, but you must then change database driver) and pass it to this lines.

```
dsn := "user=postgres password=s197328645S! dbname=todo port=5432 sslmode=disable TimeZone=Europe/Warsaw"
database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

3. Connect to your database and test.


#Important links

1. [Gin framework](https://github.com/gin-gonic/gin)
2. [GORM ORM](https://github.com/go-gorm/gorm)
3. Icons:
  *-Add https://thenounproject.com/term/add-task/945026/
  *-Delete https://thenounproject.com/search/?q=DELETE+TASK&i=525770
  *-Find https://thenounproject.com/search/?q=FIND+TASK&i=532103
  *-Update https://thenounproject.com/search/?q=UPDATE+TASK&i=2676042
  *-All https://thenounproject.com/term/task/3129877/
4. PostgreSQL https://www.postgresql.org/
