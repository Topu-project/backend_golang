package repository

import "backend_golang/domain"

type RecruitmentORM struct {
	db ORM
}

func NewRecruitmentORM(db ORM) domain.RecruitmentRepository {
	return &RecruitmentORM{db: db}
}

func (r *RecruitmentORM) Create(recruitment domain.Recruitment) error {
	record := recruitment.ToRecord()
	return r.db.Create(record)
}

func (r *RecruitmentORM) FindAll() ([]domain.Recruitment, error) {
	// ToDomain or ToEntity

	//TODO implement me
	panic("implement me")
}

func (r *RecruitmentORM) FindByID(recruitmentID int) (domain.Recruitment, error) {
	//TODO implement me
	panic("implement me")
}
