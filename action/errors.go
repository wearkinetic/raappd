package action

import "github.com/hippoai/goerr"

// ErrUnparsableJWT
func ErrUnparsableJWT() error {
	return goerr.NewS(ERR_UNPARSABLE_JWT)
}

// NoPublicEndpointSatisfiedByPayload
func NoPublicEndpointSatisfiedByPayload(payloadItf interface{}) error {
	return goerr.New(
		ERR_NO_PUBLIC_ENDPOINT_SATISFIED_BY_PAYLOAD,
		map[string]interface{}{
			"payload": payloadItf,
		},
	)
}

func ErrIncompleteParametersDescription(field string) error {
	return goerr.New(ERR_MISSING_PARAMETERS_DESCRIPTION, map[string]interface{}{
		"field": field,
	})
}

func ErrParsePayloadNotImplementedYet() error {
	return goerr.NewS(ERR_PARSE_PAYLOAD_NOT_IMPLEMENTED_YET)
}

func ErrParameterNotRequiredNorOptional(parameter string) error {
	return goerr.New(ERR_PARAMETER_NOT_REQUIRED_NOR_OPTIONAL, map[string]interface{}{
		"parameter": parameter,
	})
}
