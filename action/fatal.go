package action

import "log"

func FatalNoActionDescription() {
	log.Fatalf("%s", FATAL_NO_ACTION_DESCRIPTION)
}

func FatalIncompleteParametersDescription(resource, verb, name, field string) {
	log.Fatalf("%s - Resource: %s - Verb: %s - Name: %s - Field: %s", FATAL_MISSING_PARAMETERS_DESCRIPTION, resource, verb, name, field)
}
