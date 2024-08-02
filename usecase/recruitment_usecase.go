package usecase

import (
	"backend_golang/domain"
	"backend_golang/usecase/input"
)

type RecruitmentUsecase interface {
	Create(input input.CreateRecruitmentInput) error
}

type recruitmentUsecase struct {
	rr domain.RecruitmentRepository
}

func NewRecruitmentUsecase(rr domain.RecruitmentRepository) recruitmentUsecase {
	return recruitmentUsecase{rr: rr}
}

func (r *recruitmentUsecase) Create(input input.CreateRecruitmentInput) error {

	recruitment := domain.NewRecruitment(
		//"uuid",
		domain.PROJECT,
		//domain.ALL,
		//"#golang",
		//"#backend",
		//3,
		//1,
		//time.Now(),
		//"test@test.com",
		//"test test",
		//"test content",
	)
	err := r.rr.Create(recruitment)
	if err != nil {
	}

	return nil
}
