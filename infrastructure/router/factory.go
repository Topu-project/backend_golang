package router

import (
	"backend_golang/adapter/repository"
	"errors"
)

var (
	ErrInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGin = iota
	InstanceEcho
)

type Port int64

type Server interface {
	Listen()
}

func NewWebServerFactory(instance int, db repository.ORM, port Port) (Server, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(port, db), nil
	default:
		return nil, ErrInvalidWebServerInstance
	}
}
