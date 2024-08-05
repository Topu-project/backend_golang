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

func (m mysqlHandler) QueryRow(query string, args ...any) repository.Row {
	row := m.db.QueryRow(query, args...)
	return newMySQLRow(row)
}

func (m mysqlHandler) Query(query string, args ...any) (repository.Rows, error) {
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	row := newMySQLRows(rows)

	return &row, nil
}

type mySQLRow struct {
	row *sql.Row
}

func (m *mySQLRow) Scan(dest ...any) error {
	if err := m.row.Scan(dest...); err != nil {
		return err
	}
	return nil
}

type mySQLRows struct {
	rows *sql.Rows
}

func (m *mySQLRows) Scan(dest ...any) error {
	if err := m.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

func (m *mySQLRows) Next() bool {
	return m.rows.Next()
}

func (m *mySQLRows) Err() error {
	return m.rows.Err()
}

func (m *mySQLRows) Close() error {
	return m.rows.Close()
}

func newMySQLRows(rows *sql.Rows) mySQLRows {
	return mySQLRows{rows: rows}
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

func (m mysqlHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := m.db.QueryRowContext(ctx, query, args)
	return newMySQLRow(row)
}

func newMySQLRow(row *sql.Row) repository.Row {
	return &mySQLRow{row: row}
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
		"%s:%s@/%s?parseTime=true",
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
