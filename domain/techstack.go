package domain

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var ErrTechStackNotFound = errors.New("技術スタックが存在しません")

type (
	TechStack struct {
		id            uint
		createdAt     time.Time
		updatedAt     time.Time
		techStackName string
		recruitmentID uint
	}

	TechStackRecord struct {
		gorm.Model
		TechStackName string `gorm:"unique"`
		RecruitmentID uint   `gorm:"not null"`
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
		TechStackName: t.techStackName,
		RecruitmentID: t.recruitmentID,
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
		recruitmentID: tr.RecruitmentID,
	}
}
