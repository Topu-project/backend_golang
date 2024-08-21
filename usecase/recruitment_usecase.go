package usecase

import (
	"backend_golang/domain"
	"backend_golang/usecase/input"
	"backend_golang/usecase/output"
	"errors"
	"time"
)

type RecruitmentUsecase interface {
	Create(input input.CreateRecruitmentInput) error
	FindAll() ([]output.RecruitmentOutput, error)
	FindByID(recruitmentID int) (output.RecruitmentOutput, error)
}

type recruitmentUsecase struct {
	rr domain.RecruitmentRepository
	tr domain.TechStackRepository
	p  RecruitmentPresenter
}

func NewRecruitmentUsecase(rr domain.RecruitmentRepository, tr domain.TechStackRepository, p RecruitmentPresenter) RecruitmentUsecase {
	return &recruitmentUsecase{rr: rr, tr: tr, p: p}
}

func (r *recruitmentUsecase) FindByID(recruitmentID int) (output.RecruitmentOutput, error) {
	recruitment, err := r.rr.FindByID(recruitmentID)
	if err != nil {
		return r.p.Output(domain.Recruitment{}), err
	}
	return r.p.Output(recruitment), nil
}

func (r *recruitmentUsecase) FindAll() ([]output.RecruitmentOutput, error) {
	recruitments, err := r.rr.FindAll()
	if err != nil {
		return r.p.FindAllOutput([]domain.Recruitment{}), err
	}

	return r.p.FindAllOutput(recruitments), nil
}

func (r *recruitmentUsecase) Create(input input.CreateRecruitmentInput) error {
	var techStacks []domain.TechStack
	for _, ts := range input.TechStacks {
		techStack, err := r.tr.FindByTechStackName(ts)
		if err != nil && errors.Is(err, domain.ErrTechStackNotFound) {
			techStacks = append(techStacks, domain.NewTechStack(ts))
			continue
		}
		techStacks = append(techStacks, techStack)
	}
	//var recruitmentTechStacks []domain.RecruitmentTechStack
	//for _, ts := range input.TechStacks {
	//	techStack, err := r.tr.FindByTechStackName(ts)
	//	if err != nil && errors.Is(err, domain.ErrTechStackNotFound) {
	//		recruitmentTechStack := domain.NewRecruitmentTechStack(
	//			0,
	//			time.Now(),
	//			time.Now(),
	//			domain.NewTechStack(ts),
	//		)
	//		recruitmentTechStacks = append(recruitmentTechStacks, recruitmentTechStack)
	//		continue
	//	}
	//	recruitmentTechStacks = append(recruitmentTechStacks, domain.NewRecruitmentTechStack(
	//		0,
	//		time.Now(),
	//		time.Now(),
	//		techStack,
	//	))
	//}

	recruitment := domain.NewRecruitment(
		1,
		time.Now(),
		time.Now(),
		input.RecruitmentCategories,
		input.ProgressMethods,
		techStacks,
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
