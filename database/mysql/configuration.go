package mysql

import (
	"errors"
	"fmt"
)

// string with a function to check whether it's empty
type estring string

// isEmpty check whether given string is empty
func (es estring) isEmpty() bool {
	return len(es) == 0
}

// Config represents the data used to establish database connection
type Config struct {
	Host         estring
	Port         uint16
	DatabaseName estring
	Username     estring
	Password     estring
}

// String create the MySQL DNS
func (config Config) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
}

// validate checks whether the config is valid to be used to establish DB connection
func (config Config) validate() error {
	if config.Host.isEmpty() {
		return errors.New("host is missing in Config")
	}
	if config.Port < 1 {
		return errors.New("port is incorrectly set in Config")
	}
	if config.Username.isEmpty() || config.Password.isEmpty() {
		return errors.New("username or password is missing in Config")
	}
	return nil
}
