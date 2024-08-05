package input

import (
	"backend_golang/domain"
	"time"
)

type CreateRecruitmentInput struct {
	RecruitmentCategories domain.RecruitmentCategories `json:"recruitment_categories"`
	ProgressMethods       domain.ProgressMethods       `json:"progress_methods"`
	TechStacks            string                       `json:"tech_stacks"`
	Positions             string                       `json:"positions"`
	NumberOfPeople        int16                        `json:"number_of_people"`
	ProgressPeriod        int16                        `json:"progress_period"`
	RecruitmentDeadline   RecruitmentDeadline          `json:"recruitment_deadline"`
	Contract              string                       `json:"contract"`
	Subject               string                       `json:"subject"`
	Content               string                       `json:"content"`
}

type RecruitmentDeadline time.Time

func (r *RecruitmentDeadline) UnmarshalJSON(b []byte) error {
	// Remove the quotes from the JSON string
	s := string(b)
	s = s[1 : len(s)-1] // remove surrounding quotes

	// Parse the time string (adjust format as needed, e.g., "2006-01-02" for date only)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*r = RecruitmentDeadline(t)
	return nil
}
