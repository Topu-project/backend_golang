package router

import (
	"backend_golang/adapter/controller"
	"backend_golang/adapter/presenter"
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

	recruitmentRouter := r.Group("/recruitments")
	recruitmentSQL := repository.NewRecruitmentSQL(g.db)
	recruitmentUsecase := usecase.NewRecruitmentUsecase(recruitmentSQL, presenter.NewRecruitmentPresenter())
	recruitmentsController := controller.NewRecruitmentsController(recruitmentUsecase)

	recruitmentRouter.POST("", g.buildCreateRecruitmentController(recruitmentsController))
	recruitmentRouter.GET("", g.buildFindAllRecruitmentController(recruitmentsController))
	recruitmentRouter.GET("/:recruitment_id", g.buildFindByIDRecruitmentController(recruitmentsController))
}

func (g *ginEngine) buildCreateRecruitmentController(recruitmentsController controller.RecruitmentsController) gin.HandlerFunc {
	return func(c *gin.Context) {
		recruitmentsController.Create(c.Writer, c.Request)
	}
}

func (g *ginEngine) buildFindAllRecruitmentController(recruitmentsController controller.RecruitmentsController) gin.HandlerFunc {
	return func(c *gin.Context) {
		recruitmentsController.FindAll(c.Writer, c.Request)
	}
}

func (g *ginEngine) buildFindByIDRecruitmentController(recruitmentController controller.RecruitmentsController) gin.HandlerFunc {
	return func(c *gin.Context) {

		q := c.Request.URL.Query()
		q.Add("recruitment_id", c.Param("recruitment_id"))
		c.Request.URL.RawQuery = q.Encode()

		recruitmentController.FindByID(c.Writer, c.Request)
	}
}
