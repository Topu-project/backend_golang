package domain

import (
	"gorm.io/gorm"
	"time"
	_ "time"
)

type RecruitmentCategories string

const (
	PROJECT = RecruitmentCategories("PROJECT")
	STUDY   = RecruitmentCategories("STUDY")
)

type ProgressMethods string

const (
	ALL     = ProgressMethods("ALL")
	ONLINE  = ProgressMethods("ONLINE")
	OFFLINE = ProgressMethods("OFFLINE")
)

type (
	Recruitment struct {
		ID                    uint
		CreatedAt             time.Time
		UpdatedAt             time.Time
		DeletedAt             time.Time
		recruitmentCategories RecruitmentCategories
		progressMethods       ProgressMethods
		techStacks            string
		positions             string
		numberOfPeople        int16
		progressPeriod        int16
		recruitmentDeadline   time.Time
		contract              string
		subject               string
		content               string
	}

	RecruitmentRecord struct {
		gorm.Model
		RecruitmentCategories RecruitmentCategories
		ProgressMethods       ProgressMethods
		TechStacks            string
		Positions             string
		NumberOfPeople        int16
		ProgressPeriod        int16
		RecruitmentDeadline   time.Time
		Contract              string
		Subject               string
		Content               string
	}

	RecruitmentRepository interface {
		Create(recruitment Recruitment) error
		FindAll() ([]Recruitment, error)
		FindByID(recruitmentID int) (Recruitment, error)
	}
)

func NewRecruitment(
	recruitmentCategories RecruitmentCategories,
	progressMethods ProgressMethods,
	techStacks string,
	positions string,
	numberOfPeople int16,
	progressPeriod int16,
	recruitmentDeadline time.Time,
	contract string,
	subject string,
	content string,
) Recruitment {
	return Recruitment{
		recruitmentCategories: recruitmentCategories,
		progressMethods:       progressMethods,
		techStacks:            techStacks,
		positions:             positions,
		numberOfPeople:        numberOfPeople,
		progressPeriod:        progressPeriod,
		recruitmentDeadline:   recruitmentDeadline,
		contract:              contract,
		subject:               subject,
		content:               content,
	}
}

func (r *Recruitment) ToRecord() *RecruitmentRecord {
	return &RecruitmentRecord{
		RecruitmentCategories: r.recruitmentCategories,
		ProgressMethods:       r.progressMethods,
		TechStacks:            r.techStacks,
		Positions:             r.positions,
		NumberOfPeople:        r.numberOfPeople,
		ProgressPeriod:        r.progressPeriod,
		RecruitmentDeadline:   r.recruitmentDeadline,
		Contract:              r.contract,
		Subject:               r.subject,
		Content:               r.content,
	}
}

func (i *RecruitmentRecord) TableName() string {
	return "recruitment"
}
