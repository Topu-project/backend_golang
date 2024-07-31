package domain

type (
	Recruitment struct{}

	RecruitmentRepository interface {
		Create(recruitment Recruitment) error
	}
)

func NewRecruitment() Recruitment {
	return Recruitment{}
}
