package database

import (
	"backend_golang/adapter/repository"
	"errors"
)

var ErrInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")

const (
	InstanceMySQL = iota
)

func NewDatabaseSQLFactory(instance int) (repository.SQL, error) {
	switch instance {
	case InstanceMySQL:
		return NewMySQLHandler(newConfigMySQL())
	default:
		return nil, ErrInvalidSQLDatabaseInstance
	}
}
