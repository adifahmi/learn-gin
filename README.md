# Go rest-api
For learning purpose

- REST API with [Gin Framework](https://gin-gonic.com/)
- MariaDB/MySQL intergration with [GORM](http://gorm.io/)


#### Create Database / Account
```sql
CREATE DATABASE `gin_test` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

#### Project Setup & Run

```shell
$ go get
$ go run main.go
```

#### To explicitly compile the code before you run the server:

```shell
$ go build main.go
$ ./main
```

### Live-reload for development 

Install gin live-reload and run

```shell
$ go get github.com/codegangsta/gin
```
```shell
$ bash scripts/run-dev.sh
```

#### Postman Docs:
https://documenter.getpostman.com/view/620074/SVYtLx7M?version=latest

#### Also using:
- https://github.com/dgrijalva/jwt-go (for JWT)
- https://github.com/joho/godotenv (for env management)
- https://github.com/go-sql-driver/mysql (MySQL driver)
- https://github.com/codegangsta/gin (for live reload)
