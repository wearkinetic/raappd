package resource

import "github.com/hippoai/goerr"

func ErrWrongPayload(defaultPayloadItf interface{}) error {
	return goerr.New(
		ERR_WRONG_PAYLOAD,
		map[string]interface{}{
			"expectedPayloadStructure": defaultPayloadItf,
		},
	)
}

func ErrNoGetActionFound() error {
	return goerr.NewS(ERR_NO_GET_ACTION_FOUND)
}

func ErrMultipleGetsCanSatisfyThisQuery(getsNames []string) error {
	return goerr.New(ERR_MULTIPLE_GETS_CAN_SATISFY_THIS_QUERY, map[string]interface{}{
		"getsNames": getsNames,
	})
}
