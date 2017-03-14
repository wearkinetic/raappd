package action

const (
	ERR_NO_PUBLIC_ENDPOINT_SATISFIED_BY_PAYLOAD = "ERR_NO_PUBLIC_ENDPOINT_SATISFIED_BY_PAYLOAD"
	ERR_NO_HANDLER                              = "ERR_NO_HANDLER"
	ENV_JWT_SECRET                              = "JWT_SECRET"

	ERR_UNPARSABLE_JWT = "ERR_UNPARSABLE_JWT"
	ERR_INVALID_CLAIMS = "ERR_INVALID_CLAIMS"
	ERR_WRONG_PAYLOAD  = "ERR_WRONG_PAYLOAD"

	FATAL_NO_ACTION_DESCRIPTION          = "NO_ACTION_DESCRIPTION"
	FATAL_MISSING_PARAMETERS_DESCRIPTION = "MISSING_PARAMETERS_DESCRIPTION"
)

var jwt_secret []byte
