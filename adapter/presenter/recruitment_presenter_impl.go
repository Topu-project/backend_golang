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

func (r *recruitmentPresenter) Output(recruitment domain.Recruitment) output.RecruitmentOutput {
	readOnly := recruitment.ToReadOnly()
	return output.RecruitmentOutput{
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
	}
}

func (r *recruitmentPresenter) FindAllOutput(recruitments []domain.Recruitment) []output.RecruitmentOutput {
	outputs := make([]output.RecruitmentOutput, 0)
	for _, recruitment := range recruitments {
		readOnly := recruitment.ToReadOnly()
		outputs = append(outputs, output.RecruitmentOutput{
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
