package main

import (
	"github.com/allen-utec/vota-api/src/infrastructure"
)

func main() {
	/* Open Database Connection */
	dbConn, err := infrastructure.OpenMysqlConnection()
	if err != nil {
		return
	}

	/* Init Repositories */
	infrastructure.InitRepositories(dbConn)

	/* Init Api */
	infrastructure.InitApiRoutes()
}
