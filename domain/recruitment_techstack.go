package domain

import (
	"time"
)

type (
	RecruitmentTechStack struct {
		id            uint
		createdAt     time.Time
		updatedAt     time.Time
		techStack     TechStack
		recruitmentID uint
		techStackID   uint
	}

	RecruitmentTechStackRecord struct {
		//gorm.Model
		CreatedAt     time.Time
		UpdatedAt     time.Time
		RecruitmentID uint `gorm:"primaryKey"`
		TechStackID   uint `gorm:"primaryKey"`
	}

	RecruitmentTechStackRepository interface {
		FindByRecruitmentID(id uint) (RecruitmentTechStack, error)
	}
)

func NewRecruitmentTechStack(
	id uint,
	createdAt time.Time,
	updatedAt time.Time,
	techStack TechStack,
) RecruitmentTechStack {
	return RecruitmentTechStack{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		techStack: techStack,
	}
}

func (r *RecruitmentTechStackRecord) TableName() string {
	return "recruitment_tech_stack"
}

func (rts *RecruitmentTechStack) ToCommandRecord() RecruitmentTechStackRecord {

	return RecruitmentTechStackRecord{
		//Model:         gorm.Model{},
		RecruitmentID: 0,
		TechStackID:   0,
	}
}
