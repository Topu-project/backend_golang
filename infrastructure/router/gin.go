package router

import (
	"backend_golang/adapter/controller"
	"backend_golang/adapter/repository"
	"backend_golang/usecase"
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

	//recruitmentSQL := repository.NewRecruitmentSQL(g.db)
	//recruitmentUsecase := usecase.NewRecruitmentUsecase(&recruitmentSQL)
	//recruitmentsController := controller.NewRecruitmentsController(&recruitmentUsecase)

	//r.POST("/recruitments", g.buildCreateRecruitmentController(recruitmentsController))
	r.POST("/recruitments", g.buildCreateRecruitmentController())
}

// func (g *ginEngine) buildCreateRecruitmentController(controller controller.RecruitmentsController) gin.HandlerFunc {
func (g *ginEngine) buildCreateRecruitmentController() gin.HandlerFunc {
	return func(c *gin.Context) {
		recruitmentSQL := repository.NewRecruitmentSQL(g.db)
		recruitmentUsecase := usecase.NewRecruitmentUsecase(&recruitmentSQL)
		recruitmentsController := controller.NewRecruitmentsController(&recruitmentUsecase)
		recruitmentsController.Create(c.Writer, c.Request)
		//controller.Create(c.Writer, c.Request)
	}
}

//func (g *ginEngine) buildFindAllRecruitmentsController(controller controller.RecruitmentsController) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		controller.FindAll(c.Writer, c.Request)
//	}
//}
