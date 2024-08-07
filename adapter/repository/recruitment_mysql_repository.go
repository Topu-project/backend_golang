package repository

import (
	"backend_golang/domain"
	"errors"
	"log"
	"time"
)

type RecruitmentSQL struct {
	db SQL
}

func NewRecruitmentSQL(db SQL) domain.RecruitmentRepository {
	return &RecruitmentSQL{db: db}
}

func (r *RecruitmentSQL) FindByID(recruitmentID int) (domain.Recruitment, error) {
	var (
		recruitmentCategories domain.RecruitmentCategories
		progressMethods       domain.ProgressMethods
		techStacks            string
		positions             string
		numberOfPeople        int16
		progressPeriod        int16
		recruitmentDeadline   time.Time
		contract              string
		subject               string
		content               string
	)

	query := "SELECT * FROM recruitment WHERE id = ?"
	err := r.db.QueryRow(query, recruitmentID).Scan(
		&recruitmentCategories,
		&progressMethods,
		&techStacks,
		&positions,
		&numberOfPeople,
		&progressPeriod,
		&recruitmentDeadline,
		&contract,
		&subject,
		&content)
	if err != nil {
		return domain.Recruitment{}, err
	}
	return domain.NewRecruitment(
		recruitmentCategories,
		progressMethods,
		techStacks,
		positions,
		numberOfPeople,
		progressPeriod,
		recruitmentDeadline,
		contract,
		subject,
		content,
	), nil
}

func (r *RecruitmentSQL) FindAll() ([]domain.Recruitment, error) {

	query := `SELECT * FROM recruitment`
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.Recruitment{}, errors.Join(err, errors.New("error listing accounts"))
	}

	var recruitments = make([]domain.Recruitment, 0)
	for rows.Next() {
		var (
			id                    int
			recruitmentCategories domain.RecruitmentCategories
			progressMethods       domain.ProgressMethods
			techStacks            string
			positions             string
			numberOfPeople        int16
			progressPeriod        int16
			recruitmentDeadline   time.Time
			contract              string
			subject               string
			content               string
		)

		if err = rows.Scan(
			&id,
			&recruitmentCategories,
			&progressMethods,
			&techStacks,
			&positions,
			&numberOfPeople,
			&progressPeriod,
			&recruitmentDeadline,
			&contract,
			&subject,
			&content,
		); err != nil {
			return []domain.Recruitment{}, errors.Join(err, errors.New("error listing accounts"))
		}

		recruitments = append(recruitments, domain.NewRecruitment(
			recruitmentCategories,
			progressMethods,
			techStacks,
			positions,
			numberOfPeople,
			progressPeriod,
			recruitmentDeadline,
			contract,
			subject,
			content,
		))
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return []domain.Recruitment{}, err
	}

	return recruitments, nil
}

func (r *RecruitmentSQL) Create(recruitment domain.Recruitment) error {
	rd := recruitment.ToRecord()
	var query = `INSERT INTO recruitment(recruitment_categories,
										 progress_methods,
										 tech_stacks,
										 positions,
										 number_of_people,
										 progress_period,
										 recruitment_deadline,
										 contract,
										 subject,
										 content)
				 VALUES (?,?,?,?,?,?,?,?,?,?)`

	err := r.db.ExecuteContext(query,
		rd.RecruitmentCategories,
		rd.ProgressMethods,
		rd.TechStacks,
		rd.Positions,
		rd.NumberOfPeople,
		rd.ProgressPeriod,
		rd.RecruitmentDeadline,
		rd.Contract,
		rd.Subject,
		rd.Content,
	)
	if err != nil {
		log.Println("[recruitment_mysql_repository] : ", err)
		return err
	}
	return err
}
