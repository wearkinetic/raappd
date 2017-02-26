package resource

import "log"

func FatalNoResourceDescription() {
	log.Fatalf("%s", FATAL_NO_RESOURCE_DESCRIPTION)
}

func FatalNoResourceEndpoint() {
	log.Fatalf("%s", FATAL_NO_RESOURCE_ENDPOINT)
}

func FatalGetActionAlreadyExists(actionName string) {
	log.Fatalf("%s - Action name: %s", FATAL_GET_ACTION_ALREADY_EXISTS, actionName)
}
