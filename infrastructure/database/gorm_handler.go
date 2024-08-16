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

func (g *gormHandler) Find(dst any, cond ...any) error {
	result := g.db.Find(dst, cond...)
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
