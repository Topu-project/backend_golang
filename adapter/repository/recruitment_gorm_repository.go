package repository

import (
	"backend_golang/domain"
	"errors"
	"gorm.io/gorm"
)

type RecruitmentORM struct {
	db ORM
}

func NewRecruitmentORM(db ORM) domain.RecruitmentRepository {
	return &RecruitmentORM{db: db}
}

func (r *RecruitmentORM) Create(recruitment domain.Recruitment) error {
	record := recruitment.ToCommandRecord()
	return r.db.Create(record)
}

func (r *RecruitmentORM) FindAll() ([]domain.Recruitment, error) {
	var recruitmentRecords []domain.RecruitmentRecord
	//if err := r.db.Find(&recruitmentRecords); err != nil {
	if err := r.db.FindWithPreload(&recruitmentRecords, "TechStacks"); err != nil {
		return nil, err
	}

	var recruitments []domain.Recruitment
	for _, record := range recruitmentRecords {
		recruitments = append(recruitments, record.ToDomain())
	}

	return recruitments, nil
}

func (r *RecruitmentORM) FindByID(recruitmentID int) (domain.Recruitment, error) {
	var record domain.RecruitmentRecord
	//if err := r.db.Find(&record, recruitmentID); err != nil {
	if err := r.db.FindWithPreload(&record, "TechStacks", recruitmentID); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return domain.Recruitment{}, domain.ErrRecruitmentNotFound
		default:
			return domain.Recruitment{}, err
		}
	}

	return record.ToDomain(), nil
}
