# TODO

# Simple TODO app using Postgres + GORM + Gin framework

1. Go to [Go website](https://golang.org/) and set up correct paths.

2. Configure PostgreSQL (or other DB, but you must then change database driver) and pass it to this lines.

```
dsn := "user=postgres password=s197328645S! dbname=todo port=5432 sslmode=disable TimeZone=Europe/Warsaw"
database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

3. Connect to your database and test.
