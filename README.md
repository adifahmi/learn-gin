# Go rest-api
For learning purpose

- REST API with [Gin Framework](https://gin-gonic.com/)
- MariaDB/MySQL intergration with [GORM](http://gorm.io/)


#### Create Database / Account
```sql
CREATE DATABASE gin_test;
```

#### Project Setup & Run

```
$ go get
$ go run main.go
```

#### To explicitly compile the code before you run the server:

```
$ go build main.go
$ ./main
```

### Live-reload for development 
```
$ scripts/run-dev.sh
```

#### Postman Docs:
https://documenter.getpostman.com/view/620074/SVYtLx7M?version=latest

#### Also using:
- https://github.com/dgrijalva/jwt-go (for JWT)
- https://github.com/joho/godotenv (for env management)
- https://github.com/go-sql-driver/mysql (MySQL driver)
- https://github.com/codegangsta/gin (for live reload)