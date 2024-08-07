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

func (g *gormHandler) AutoMigrate(dst ...any) {
	if err := g.db.AutoMigrate(dst...); err != nil {
		panic("failed to auto migrate database")
	}
}

func NewGormHandler(c *config) repository.ORM {
	dsn := fmt.Sprintf(
		"%s:%s@/%s?parseTime=true",
		c.user,
		c.password,
		c.database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &gormHandler{db: db}
}

func (g *gormHandler) Create(value any) error {
	result := g.db.Create(value)
	if result.Error != nil {
		log.Println("[gormHandler Create] error:]", result.Error)
		return result.Error
	}
	return nil
}
