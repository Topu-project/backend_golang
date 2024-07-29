package infrastructure

import (
	"backend_golang/adapter/repository"
	"backend_golang/infrastructure/database"
	"log"
)

type config struct {
	dbSQL repository.SQL
}

func NewConfig() *config {
	return &config{}
}

func (c *config) DbSQL(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		log.Fatalln(err, "Could not make a connection to database")
	}

	log.Println("Successfully connected to the SQL database")
	c.dbSQL = db
	return c
}
