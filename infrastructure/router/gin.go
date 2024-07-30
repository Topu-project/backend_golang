package router

import (
	"backend_golang/adapter/repository"
	"log"
)
import "github.com/gin-gonic/gin"

type ginEngine struct {
	router *gin.Engine
	port   Port
	db     repository.SQL
}

func newGinServer(port Port, db repository.SQL) *ginEngine {
	return &ginEngine{
		router: gin.New(),
		port:   port,
		db:     db,
	}
}

func (g *ginEngine) Listen() {
	log.Println("starting server...")
}
