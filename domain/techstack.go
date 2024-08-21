package domain

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var ErrTechStackNotFound = errors.New("技術スタックが存在しません")

type (
	TechStack struct {
		id                    uint
		createdAt             time.Time
		updatedAt             time.Time
		techStackName         string
		RecruitmentTechStacks []RecruitmentTechStack
	}

	TechStackRecord struct {
		gorm.Model
		TechStackName string `gorm:"unique"`
		//RecruitmentTechStacks []RecruitmentTechStackRecord `gorm:"foreignKey:TechStackID"`
		Recruitments []RecruitmentRecord `gorm:"many2many:recruitment_tech_stack"`
	}

	TechStackRepository interface {
		FindByTechStackName(techStackName string) (TechStack, error)
	}
)

func NewTechStack(techStackName string) TechStack {
	return TechStack{techStackName: techStackName}
}

func (t *TechStack) ToCommandRecord() TechStackRecord {
	return TechStackRecord{
		Model: gorm.Model{
			ID:        t.id,
			CreatedAt: t.createdAt,
			UpdatedAt: t.updatedAt,
		},
		TechStackName: t.techStackName,
	}
}

func (t *TechStack) ToReadRecord() TechStackRecord {
	return TechStackRecord{
		Model: gorm.Model{
			ID:        t.id,
			CreatedAt: t.createdAt,
			UpdatedAt: t.updatedAt,
		},
		TechStackName: t.techStackName,
	}
}

func (tr *TechStackRecord) TableName() string {
	return "tech_stack"
}

func (tr *TechStackRecord) ToDomain() TechStack {
	return TechStack{
		id:            tr.ID,
		createdAt:     tr.CreatedAt,
		updatedAt:     tr.UpdatedAt,
		techStackName: tr.TechStackName,
		//recruitmentID: tr.RecruitmentID,
	}
}
