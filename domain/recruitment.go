package domain

import (
	"errors"
	"gorm.io/gorm"
	"time"
	_ "time"
)

var ErrRecruitmentNotFound = errors.New("応募ポストが存在しません")

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
		id                    uint
		createdAt             time.Time
		updatedAt             time.Time
		recruitmentCategories RecruitmentCategories
		progressMethods       ProgressMethods
		techStacks            []TechStack
		positions             string
		numberOfPeople        int16
		progressPeriod        int16
		recruitmentDeadline   time.Time
		contract              string
		subject               string
		content               string
	}

	//RecruitmentRecord gormはフィールドがpublicである必要があるので別途Record構造体をよいしておく
	RecruitmentRecord struct {
		gorm.Model
		RecruitmentCategories RecruitmentCategories
		ProgressMethods       ProgressMethods
		TechStacks            []TechStackRecord `gorm:"many2many:recruitment_tech_stack"`
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
	id uint,
	createdAt time.Time,
	updatedAt time.Time,
	recruitmentCategories RecruitmentCategories,
	progressMethods ProgressMethods,
	techStacks []TechStack,
	positions string,
	numberOfPeople int16,
	progressPeriod int16,
	recruitmentDeadline time.Time,
	contract string,
	subject string,
	content string,
) Recruitment {
	return Recruitment{
		id:                    id,
		createdAt:             createdAt,
		updatedAt:             updatedAt,
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

func (r *Recruitment) ToReadRecord() *RecruitmentRecord {
	//var techStackRecords []TechStackRecord
	//for _, ts := range r.techStacks {
	//	techStackRecords = append(techStackRecords, ts.ToCommandRecord())
	//}

	return &RecruitmentRecord{
		Model: gorm.Model{
			ID:        r.id,
			CreatedAt: r.createdAt,
			UpdatedAt: r.updatedAt,
		},
		RecruitmentCategories: r.recruitmentCategories,
		ProgressMethods:       r.progressMethods,
		//TechStacks:            techStackRecords,
		Positions:           r.positions,
		NumberOfPeople:      r.numberOfPeople,
		ProgressPeriod:      r.progressPeriod,
		RecruitmentDeadline: r.recruitmentDeadline,
		Contract:            r.contract,
		Subject:             r.subject,
		Content:             r.content,
	}
}

func (r *Recruitment) ToCommandRecord() *RecruitmentRecord {
	//var records []RecruitmentTechStackRecord
	//for _, rt := range r.recruitmentTechStacks {
	//	records = append(records, rt.ToCommandRecord())
	//}

	var records []TechStackRecord
	for _, ts := range r.techStacks {
		records = append(records, ts.ToCommandRecord())
	}

	return &RecruitmentRecord{
		RecruitmentCategories: r.recruitmentCategories,
		ProgressMethods:       r.progressMethods,
		TechStacks:            records,
		Positions:             r.positions,
		NumberOfPeople:        r.numberOfPeople,
		ProgressPeriod:        r.progressPeriod,
		RecruitmentDeadline:   r.recruitmentDeadline,
		Contract:              r.contract,
		Subject:               r.subject,
		Content:               r.content,
	}
}

func (r *RecruitmentRecord) ToDomain() Recruitment {
	//var techStacks []TechStack
	//for _, record := range r.TechStacks {
	//	techStacks = append(techStacks, record.ToDomain())
	//}

	return Recruitment{}
	//return NewRecruitment(
	//	r.ID,
	//	r.CreatedAt,
	//	r.UpdatedAt,
	//	r.RecruitmentCategories,
	//	r.ProgressMethods,
	//	//techStacks,
	//	r.Positions,
	//	r.NumberOfPeople,
	//	r.ProgressPeriod,
	//	r.RecruitmentDeadline,
	//	r.Contract,
	//	r.Subject,
	//	r.Content,
	//)
}

// TableName gorm で table 名を custom するためには　この関数を override する必要がある
func (r *RecruitmentRecord) TableName() string {
	return "recruitment"
}
