# go-db-shell
[![Go Reference](https://pkg.go.dev/badge/github.com/M-Arthur/go-db-shell.svg)](https://pkg.go.dev/github.com/M-Arthur/go-db-shell)
[![Go version](https://img.shields.io/badge/go-v1.16.5-blue.svg)](https://github.com/golang/go/tree/go1.16.5)
[![Go MySQL Driver version](https://img.shields.io/badge/go_mysql_driver-v1.6.0-blue.svg)](https://github.com/go-sql-driver/mysql/tree/v1.6.0)

https://github.com/go-sql-driver/mysql/tree/v1.6.0

An extra layer which make communicate with database safer. All the query with placeholders will be escaped.

### Example
Below is the example of how to create the db connection
```go
db, err := mysql.Open(mysql.Config{
    Host:         "127.0.0.1",
    Port:         3306,
    Username:     "username",
    Password:     "password",
    DatabaseName: "database",
})

if err != nil {
    panic(err)
}

defer db.Close()
```
How to execute SQL query and get results are the same as what you would do with go-mysql-driver.

### Avaiable functoins
* query()
* queryRow()
* execute()