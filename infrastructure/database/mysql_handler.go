package database

import (
	"database/sql"
	"fmt"
	"log"
)

type mysqlHandler struct {
	db *sql.DB
}

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
