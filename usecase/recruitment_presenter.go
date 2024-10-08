package usecase

import (
	"backend_golang/domain"
	"backend_golang/usecase/output"
)

type RecruitmentPresenter interface {
	FindAllOutput([]domain.Recruitment) []output.RecruitmentOutput
	Output(domain.Recruitment) output.RecruitmentOutput
}
