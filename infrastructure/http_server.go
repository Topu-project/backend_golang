package infrastructure

import (
	"backend_golang/adapter/repository"
	"backend_golang/infrastructure/database"
	"backend_golang/infrastructure/router"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type config struct {
	env           string
	dbSQL         repository.SQL
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalln("Error when loading .env file", err)
	}
	env := os.Getenv("ENV")
	log.Println(env, " IS LOADED")
	return &config{env: env}
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

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.dbSQL,
		c.webServerPort,
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully configured router server")
	c.webServer = s
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
