package controller

import (
	"backend_golang/adapter/response"
	"backend_golang/usecase"
	"backend_golang/usecase/input"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type RecruitmentsController interface {
	Create(w http.ResponseWriter, req *http.Request)
	FindAll(w http.ResponseWriter, req *http.Request)
	FindByID(w http.ResponseWriter, req *http.Request)
	Delete(w http.ResponseWriter, req *http.Request)
}

type recruitmentController struct {
	uc usecase.RecruitmentUsecase
}

func NewRecruitmentsController(uc usecase.RecruitmentUsecase) RecruitmentsController {
	return &recruitmentController{uc: uc}
}

func (r *recruitmentController) FindByID(w http.ResponseWriter, req *http.Request) {

	recruitmentID := req.URL.Query().Get("recruitment_id")
	parsedID, _ := strconv.ParseInt(recruitmentID, 10, 64)
	output, err := r.uc.FindByID(int(parsedID))
	if err != nil {
		response.NewError(http.StatusBadRequest, err).Send(w)
		return
	}

	response.NewSuccess(http.StatusOK, output).Send(w)
}

func (r *recruitmentController) Delete(w http.ResponseWriter, req *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r *recruitmentController) FindAll(w http.ResponseWriter, req *http.Request) {
	output, err := r.uc.FindAll()
	if err != nil {
		log.Println("[find_all_recruitment] : ", err.Error())
		response.NewError(http.StatusInternalServerError, err).Send(w)
		return
	}

	log.Println("[find_all_recruitment] : ", "success find all recruitments")
	response.NewSuccess(http.StatusOK, output).Send(w)
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
