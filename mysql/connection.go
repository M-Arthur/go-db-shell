package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// connection represents the DB object which is used to execute SQL
type connection interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) (*sql.Row, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Close() error
	SelectRow(interface{}, string, string) error
}

// mysql is one solid implemetation of connection interface
type mysql struct {
	db     *sql.DB
	config Config
}

// Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
func (m mysql) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	queryString, err := queryBuilder{sql: sql, parameters: args}.build()
	if err != nil {
		return nil, err
	}

	return m.db.Query(queryString)
}

// QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.
func (m mysql) QueryRow(sql string, args ...interface{}) (*sql.Row, error) {
	queryString, err := queryBuilder{sql: sql, parameters: args}.build()
	if err != nil {
		return nil, err
	}

	return m.db.QueryRow(queryString), nil
}

// Exec executes a query without returning any rows. The args are for any placeholder parameters in the query.
func (m mysql) Exec(sql string, args ...interface{}) (sql.Result, error) {
	queryString, err := queryBuilder{sql: sql, parameters: args}.build()
	if err != nil {
		return nil, err
	}

	return m.db.Exec(queryString)
}

func (m mysql) SelectRow(structure interface{}, tableName string, conditions string) error {
	rStruct := newReflectStruct(structure)
	fieldArray := make([]string, rStruct.numField)
	addressArray := make([]interface{}, rStruct.numField)
	for i := 0; i < rStruct.numField; i++ {
		fieldArray[i] = strings.ToLower(fmt.Sprintf("`%s`", rStruct.getFiledNameByIndex(i)))
		addressArray[i] = rStruct.getFieldAddressByIndex(i)
	}
	sql := fmt.Sprintf("SELECT %s FROM %s %s", strings.Join(fieldArray, ", "), tableName, conditions)

	return m.db.QueryRow(sql).Scan(addressArray...)
}

// Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.
func (m mysql) Close() error {
	return m.db.Close()
}

// Open opens a database connection to MySQL database server
func Open(config Config) (connection, error) {
	err := config.validate()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", config.String())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return mysql{
		db:     db,
		config: config,
	}, nil
}
