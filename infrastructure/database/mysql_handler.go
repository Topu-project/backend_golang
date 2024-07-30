package database

import (
	"database/sql"
	"fmt"
	"log"
)

type mysqlHandler struct {
	db *sql.DB
}

//func (m *mysqlHandler) ExecuteContext(ctx context.Context, s string, i ...interface{}) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m *mysqlHandler) QueryContext(ctx context.Context, s string, i ...interface{}) (repository.Rows, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m *mysqlHandler) QueryRowContext(ctx context.Context, s string, i ...interface{}) repository.Row {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m *mysqlHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
//	//TODO implement me
//	panic("implement me")
//}

func NewMySQLHandler(c *config) (*mysqlHandler, error) {
	ds := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
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
