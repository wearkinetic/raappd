package raappd

import "log"

func FatalResourceAlreadyExists(endpoint string) {
	log.Fatalf("%s - Endpoint: %s", FATAL_RESOURCE_ALREADY_EXISTS, endpoint)
}
