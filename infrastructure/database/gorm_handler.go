package database

import (
	"backend_golang/adapter/repository"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type gormHandler struct {
	db *gorm.DB
}

func (g *gormHandler) FindWithPreload(dst any, preload string, cond ...any) error {
	result := g.db.Preload(preload).Where("id = ?", cond...).First(dst)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	if result.RowsAffected < 1 {
		log.Println("[gormHandler.Find] : ", gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (g *gormHandler) Find(dst any, cond ...any) error {
	result := g.db.Find(dst, cond...)
	// user.UserRoles에 사용자와 연결된 UserRole들이 포함됩니다.
	// 각 UserRole에는 UserRole.Role에 해당하는 Role 정보도 포함됩니다.
	//var techStacks []domain.TechStackRecord
	//
	//result := g.db.
	//	Preload("TechStacks").
	//	Model(dst).
	//	Joins("JOIN recruitment_tech_stack ON recruitment.id = recruitment_tech_stack.recruitment_record_id").
	//	Joins("JOIN tech_stack ON recruitment_tech_stack.tech_stack_record_id = tech_stack.id").
	//	Where("recruitment.id = ?", cond...).
	//	Find(&techStacks)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	if result.RowsAffected < 1 {
		log.Println("[gormHandler.Find] : ", gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (g *gormHandler) AutoMigrate(dst ...any) {
	//g.db.SetupJoinTable(&domain.RecruitmentRecord{}, "TechStacks", &domain.RecruitmentTechStackRecord{})
	if err := g.db.AutoMigrate(dst...); err != nil {
		panic("[gormHandler.AutoMigrate] : failed to auto migrate database")
	}
}

func (g *gormHandler) Create(value any) error {
	result := g.db.Create(value)
	if result.Error != nil {
		log.Println("[gormHandler.Create] : ", result.Error)
		return result.Error
	}
	return nil
}

func newGormHandler(c *config) repository.ORM {
	//"%s:%s@/%s?parseTime=true&time_zone=Asia/Tokyo",
	dsn := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		c.user,
		c.password,
		c.database,
	)

	// あとでdsnとgorm dialectも外部から注入できるようにした方が良さそう
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[gormHandler.newGormHandler] : failed to connect database")
	}

	return &gormHandler{db: db}
}
