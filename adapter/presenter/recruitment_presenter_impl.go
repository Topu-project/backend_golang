package presenter

import (
	"backend_golang/domain"
	"backend_golang/usecase"
	"backend_golang/usecase/output"
)

type recruitmentPresenter struct {
}

func NewRecruitmentPresenter() usecase.RecruitmentPresenter {
	return &recruitmentPresenter{}
}

func (r *recruitmentPresenter) FindAllOutput(recruitments []domain.Recruitment) []output.FindAllRecruitmentOutput {
	outputs := make([]output.FindAllRecruitmentOutput, 0)
	for _, recruitment := range recruitments {
		readOnly := recruitment.ToReadOnly()
		outputs = append(outputs, output.FindAllRecruitmentOutput{
			ID:                    readOnly.ID,
			RecruitmentCategories: readOnly.RecruitmentCategories,
			ProgressMethods:       readOnly.ProgressMethods,
			TechStacks:            readOnly.TechStacks,
			Positions:             readOnly.Positions,
			NumberOfPeople:        readOnly.NumberOfPeople,
			ProgressPeriod:        readOnly.ProgressPeriod,
			RecruitmentDeadline:   readOnly.RecruitmentDeadline,
			Contract:              readOnly.Contract,
			Subject:               readOnly.Subject,
			Content:               readOnly.Content,
		})
	}
	return outputs
}
