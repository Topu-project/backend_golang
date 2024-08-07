package main

import (
	"backend_golang/domain"
	"backend_golang/infrastructure"
	"backend_golang/infrastructure/database"
	"backend_golang/infrastructure/router"
	"os"
)

func main() {
	app := infrastructure.NewConfig().
		ORM(database.InstanceGorm).
		Migrate(&domain.RecruitmentRecord{})

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin).
		Start()
}
