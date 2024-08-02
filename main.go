package main

import (
	"backend_golang/infrastructure"
	"backend_golang/infrastructure/database"
	"backend_golang/infrastructure/router"
	"os"
)

func main() {
	app := infrastructure.NewConfig().
		DbSQL(database.InstanceMySQL)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin).
		Start()
}
