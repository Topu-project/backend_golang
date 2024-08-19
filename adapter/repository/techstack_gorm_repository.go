package repository

import (
	"backend_golang/domain"
	"errors"
	"gorm.io/gorm"
)

type TechStackORM struct {
	db ORM
}

func NewTechStackORM(db ORM) domain.TechStackRepository {
	return &TechStackORM{db: db}
}

func (t *TechStackORM) FindByTechStackName(techStackName string) (domain.TechStack, error) {
	var record domain.TechStackRecord
	if err := t.db.Find(&record, "tech_stack_name = ?", techStackName); err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return domain.TechStack{}, domain.ErrTechStackNotFound
		default:
			return domain.TechStack{}, err
		}
	}
	return record.ToDomain(), nil

}
