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
	record := recruitment.ToRecord()
	return output.RecruitmentOutput{
		ID:                    record.ID,
		CreatedAt:             record.CreatedAt,
		UpdatedAt:             record.UpdatedAt,
		RecruitmentCategories: record.RecruitmentCategories,
		ProgressMethods:       record.ProgressMethods,
		TechStacks:            record.TechStacks,
		Positions:             record.Positions,
		NumberOfPeople:        record.NumberOfPeople,
		ProgressPeriod:        record.ProgressPeriod,
		RecruitmentDeadline:   record.RecruitmentDeadline,
		Contract:              record.Contract,
		Subject:               record.Subject,
		Content:               record.Content,
	}
}

func (r *recruitmentPresenter) FindAllOutput(recruitments []domain.Recruitment) []output.RecruitmentOutput {
	outputs := make([]output.RecruitmentOutput, 0)
	for _, recruitment := range recruitments {
		readOnly := recruitment.ToRecord()
		outputs = append(outputs, output.RecruitmentOutput{
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
