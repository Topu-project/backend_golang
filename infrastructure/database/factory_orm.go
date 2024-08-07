package database

import (
	"backend_golang/adapter/repository"
	"errors"
)

var ErrInvalidORMDatabaseInstance = errors.New("invalid orm db instance")

const (
	InstanceGorm = iota
)

func NewORMFactory(instance int) (repository.ORM, error) {
	switch instance {
	case InstanceGorm:
		return NewGormHandler(newConfigGorm()), nil
	default:
		return nil, ErrInvalidORMDatabaseInstance
	}
}
