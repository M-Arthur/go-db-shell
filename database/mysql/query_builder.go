package mysql

import (
	"errors"
	"fmt"
	"strings"
)

// queryBuilder represents the structure which is used to biuld safe query
type queryBuilder struct {
	sql        string
	parameters []interface{}
}

// validate checks whether the given query data is valid
func (qb queryBuilder) validate() error {
	if strings.Count(qb.sql, "?") != len(qb.parameters) {
		return errors.New("the number of identifier and parameters do not match")
	}
	return nil
}

// build creates a safe SQL query
func (qb queryBuilder) build() (string, error) {
	if err := qb.validate(); err != nil {
		return "", err
	}
	parameters := make([]interface{}, len(qb.parameters))
	for index, argument := range qb.parameters {
		parameters[index] = Escape(fmt.Sprintf("%v", argument))
	}
	sql := strings.ReplaceAll(qb.sql, "'?'", "'%s'")
	sql = strings.ReplaceAll(sql, "?", "'%s'")
	return fmt.Sprintf(sql, parameters...), nil
}
