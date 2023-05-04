package infrastructure

import (
	"github.com/allen-utec/vota-api/src/application"
)

func Bootstrap() {

	/* Open Database Connection */
	dbConn, err := OpenMysqlConnection()
	if err != nil {
		return
	}

	/* Init Repositories */
	pollRepo := &PollRepository{dbConn}
	userRepo := &UserRepository{dbConn}

	/* Init Application Services */
	application.InitServices(pollRepo, userRepo)

	/* Init Apis */
	InitApiRoutes()
}
