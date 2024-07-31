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

func NewRecruitmentUsecase(rr domain.RecruitmentRepository) RecruitmentUsecase {
	return &recruitmentUsecase{rr: rr}
}

func (r *recruitmentUsecase) Create(input input.CreateRecruitmentInput) error {

	recruitment := domain.NewRecruitment()
	err := r.rr.Create(recruitment)
	if err != nil {
	}

	return nil
}
