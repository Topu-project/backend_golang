package router

import (
	"backend_golang/adapter/controller"
	"backend_golang/adapter/repository"
	"fmt"
	"log"
	"net/http"
	"time"
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
	gin.Recovery()
	g.setAppHandlers(g.router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}

}

func (g *ginEngine) setAppHandlers(r *gin.Engine) {

	recruitmentsController := controller.NewRecruitmentsController()

	r.GET("/recruitments", g.buildFindAllRecruitmentsController(recruitmentsController))
}

func (g *ginEngine) buildFindAllRecruitmentsController(controller controller.RecruitmentsController) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller.FindAll(c.Writer, c.Request)
	}
}
