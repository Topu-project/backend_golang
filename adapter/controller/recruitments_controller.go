package controller

import (
	"backend_golang/adapter/response"
	"backend_golang/usecase"
	"backend_golang/usecase/input"
	"encoding/json"
	"log"
	"net/http"
)

type RecruitmentsController interface {
	Create(w http.ResponseWriter, req *http.Request)
}

type recruitmentController struct {
	uc usecase.RecruitmentUsecase
}

func NewRecruitmentsController(uc usecase.RecruitmentUsecase) RecruitmentsController {
	return &recruitmentController{uc: uc}
}

func (r *recruitmentController) Create(w http.ResponseWriter, req *http.Request) {

	var input input.CreateRecruitmentInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		log.Println("[create_recruitment] : ", err.Error())
		response.NewError(http.StatusBadRequest, err).Send(w)
		return
	}
	defer req.Body.Close()
	err := r.uc.Create(input)
	if err != nil {
		log.Println("[create_recruitment] : ", err.Error())
		response.NewError(http.StatusBadRequest, err).Send(w)
		return
	}
	log.Println("[create_recruitment] : ", "success creating recruitment")
	response.NewSuccess(http.StatusCreated, nil).Send(w)
}
