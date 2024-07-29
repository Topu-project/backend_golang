package backend_golang

import (
	"backend_golang/infrastructure"
	"backend_golang/infrastructure/database"
)

func main() {
	app := infrastructure.NewConfig().
		DbSQL(database.InstanceMySQL)

}
