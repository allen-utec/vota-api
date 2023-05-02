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
	infrastructure.InitUserRepository(dbConn)
	infrastructure.InitPollRepository(dbConn)

	/* Init Api */
	infrastructure.InitApiRoutes()
}
