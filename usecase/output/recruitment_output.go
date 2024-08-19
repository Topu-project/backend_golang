package output

import (
	"backend_golang/domain"
	"time"
)

type CreateRecruitmentOutput struct {
}

type RecruitmentOutput struct {
	ID                    uint                         `json:"id"`
	CreatedAt             time.Time                    `json:"created_at"`
	UpdatedAt             time.Time                    `json:"updated_at"`
	RecruitmentCategories domain.RecruitmentCategories `json:"recruitment_categories"`
	ProgressMethods       domain.ProgressMethods       `json:"progress_methods"`
	TechStacks            []string                     `json:"tech_stacks"`
	Positions             string                       `json:"positions"`
	NumberOfPeople        int16                        `json:"number_of_people"`
	ProgressPeriod        int16                        `json:"progress_period"`
	RecruitmentDeadline   time.Time                    `json:"recruitment_deadline"`
	Contract              string                       `json:"contract"`
	Subject               string                       `json:"subject"`
	Content               string                       `json:"content"`
}
