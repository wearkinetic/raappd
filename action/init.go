package action

import (
	"log"

	"github.com/hippoai/env"
)

func init() {

	// Need JWT_SECRET passed as environment variable
	parsed, err := env.Parse(
		ENV_JWT_SECRET,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}

	jwt_secret = []byte(parsed[ENV_JWT_SECRET])

}
