package usecase

import (
	"backend_golang/domain"
	"backend_golang/usecase/input"
	"backend_golang/usecase/output"
	"time"
)

type RecruitmentUsecase interface {
	Create(input input.CreateRecruitmentInput) error
	FindAll() ([]output.FindAllRecruitmentOutput, error)
}

type recruitmentUsecase struct {
	rr domain.RecruitmentRepository
	p  RecruitmentPresenter
}

func NewRecruitmentUsecase(rr domain.RecruitmentRepository, p RecruitmentPresenter) RecruitmentUsecase {
	return &recruitmentUsecase{rr: rr, p: p}
}

func (r *recruitmentUsecase) FindAll() ([]output.FindAllRecruitmentOutput, error) {
	recruitments, err := r.rr.FindAll()
	if err != nil {
		return r.p.FindAllOutput([]domain.Recruitment{}), err
	}

	return r.p.FindAllOutput(recruitments), nil
}

func (r *recruitmentUsecase) Create(input input.CreateRecruitmentInput) error {

	recruitment := domain.NewRecruitment(
		nil,
		input.RecruitmentCategories,
		input.ProgressMethods,
		input.TechStacks,
		input.Positions,
		input.NumberOfPeople,
		input.ProgressPeriod,
		time.Time(input.RecruitmentDeadline),
		input.Contract,
		input.Subject,
		input.Content,
	)
	err := r.rr.Create(recruitment)
	if err != nil {
		return err
	}

	return nil
}
