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
	record := recruitment.ToReadRecord()

	var techStacks []string
	for _, techStack := range record.TechStacks {
		techStacks = append(techStacks, techStack.TechStackName)
	}
	return output.RecruitmentOutput{
		ID:                    record.ID,
		CreatedAt:             record.CreatedAt,
		UpdatedAt:             record.UpdatedAt,
		RecruitmentCategories: record.RecruitmentCategories,
		ProgressMethods:       record.ProgressMethods,
		TechStacks:            techStacks,
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
		record := recruitment.ToReadRecord()
		var techStacks []string
		for _, techStack := range record.TechStacks {
			techStacks = append(techStacks, techStack.TechStackName)
		}
		outputs = append(outputs, output.RecruitmentOutput{
			ID:                    record.ID,
			CreatedAt:             record.CreatedAt,
			UpdatedAt:             record.UpdatedAt,
			RecruitmentCategories: record.RecruitmentCategories,
			ProgressMethods:       record.ProgressMethods,
			TechStacks:            techStacks,
			Positions:             record.Positions,
			NumberOfPeople:        record.NumberOfPeople,
			ProgressPeriod:        record.ProgressPeriod,
			RecruitmentDeadline:   record.RecruitmentDeadline,
			Contract:              record.Contract,
			Subject:               record.Subject,
			Content:               record.Content,
		})
	}
	return outputs
}
