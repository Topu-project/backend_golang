package repository

type ORM interface {
	AutoMigrate(dst ...any)
	Create(value any) error
	Find(dst any, cond ...any) error
	FindWithPreload(dst any, preload string, cond ...any) error
}
