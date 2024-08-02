package database

import (
	"backend_golang/adapter/repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type mysqlHandler struct {
	db *sql.DB
}

func (m mysqlHandler) ExecuteContext(query string, args ...interface{}) error {
	_, err := m.db.Exec(query, args...)
	if err != nil {
		log.Println("[mysql_handler] : ", err)
		return err
	}

	return nil
}

func (m mysqlHandler) QueryContext(ctx context.Context, s string, i ...interface{}) (repository.Rows, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlHandler) QueryRowContext(ctx context.Context, s string, i ...interface{}) repository.Row {
	//TODO implement me
	panic("implement me")
}

func (m mysqlHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	//TODO implement me
	panic("implement me")
}

func NewMySQLHandler(c *config) (*mysqlHandler, error) {
	ds := fmt.Sprintf(
		//db, err := sql.Open("mysql", "user:password@/dbname")
		//    db, err := sql.Open("mysql", "kazuhira:password@(172.17.0.2:3306)/practice")
		//"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		"%s:%s@/%s",
		c.user,
		c.password,
		c.database,
	)

	fmt.Println(ds)
	db, err := sql.Open(c.driver, ds)
	if err != nil {
		return &mysqlHandler{}, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &mysqlHandler{db: db}, nil
}
