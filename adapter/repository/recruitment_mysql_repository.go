package repository

import (
	"backend_golang/domain"
	"log"
)

type RecruitmentSQL struct {
	db SQL
}

func NewRecruitmentSQL(db SQL) RecruitmentSQL {
	return RecruitmentSQL{
		db: db,
	}
}

func (r *RecruitmentSQL) Create(recruitment domain.Recruitment) error {
	var query = `INSERT INTO recruitment(recruitment_categories) VALUES (?)`
	//	var query = `
	//		INSERT INTO
	//			recruitment(
	//			            recruitment_categories
	//-- 			            progressMethods,
	//-- 			            techStacks,
	//-- 			            positions,
	//-- 			            numberOfPeople,
	//-- 			            progressPeriod,
	//-- 			            recruitmentDeadline,
	//-- 			            contract,
	//-- 			            subject,
	//-- 			            content
	//	      )
	//		VALUES
	//			($1
	//-- 			 $2,
	//-- 			 $3,
	//-- 			 $4,
	//-- 			 $5,
	//-- 			 $6,
	//-- 			 $7,
	//-- 			 $8,
	//-- 			 $9,
	//-- 			 $10)
	//			)`

	err := r.db.ExecuteContext(query, recruitment.RecruitmentCategories())
	if err != nil {
		log.Println("[recruitment_mysql_repository] : ", err)
		return err
	}
	return err
}
