# go-db-shell
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