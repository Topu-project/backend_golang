package controller

import "net/http"

type RecruitmentsController interface {
	FindAll(w http.ResponseWriter, req *http.Request)
}

type recruitmentController struct {
}

func (r *recruitmentController) FindAll(w http.ResponseWriter, req *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewRecruitmentsController() RecruitmentsController {
	return &recruitmentController{}
}
