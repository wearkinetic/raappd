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

func ErrWrongPayload(defaultPayloadItf interface{}) error {
	return goerr.New(
		ERR_WRONG_PAYLOAD,
		map[string]interface{}{
			"expectedPayloadStructure": defaultPayloadItf,
		},
	)
}
