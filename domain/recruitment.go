package domain

import _ "time"

//go:generate stringer -type=RecruitmentCategories
type RecruitmentCategories int

const (
	PROJECT RecruitmentCategories = iota
	STUDY
)

//go:generate stringer -type=ProgressMethods
type ProgressMethods int

const (
	ALL ProgressMethods = iota
	ONLINE
	OFFLINE
)

type (
	Recruitment struct {
		//id                    string
		recruitmentCategories RecruitmentCategories
		//progressMethods       ProgressMethods
		//techStacks            string
		//positions             string
		//numberOfPeople        int16
		//progressPeriod        int16
		//recruitmentDeadline   time.Time
		//contract              string
		//subject               string
		//content               string
	}

	RecruitmentRepository interface {
		Create(recruitment Recruitment) error
	}
)

func NewRecruitment(
	//id string,
	recruitmentCategories RecruitmentCategories,
	// progressMethods ProgressMethods,
	// techStacks string,
	// positions string,
	// numberOfPeople int16,
	// progressPeriod int16,
	// recruitmentDeadline time.Time,
	// contract string,
	// subject string,
	// content string,
) Recruitment {
	return Recruitment{
		//id:                    id,
		recruitmentCategories: recruitmentCategories,
		//progressMethods:       progressMethods,
		//techStacks:            techStacks,
		//positions:             positions,
		//numberOfPeople:        numberOfPeople,
		//progressPeriod:        progressPeriod,
		//recruitmentDeadline:   recruitmentDeadline,
		//contract:              contract,
		//subject:               subject,
		//content:               content,
	}
}

func (r *Recruitment) RecruitmentCategories() RecruitmentCategories {
	return r.recruitmentCategories
}
