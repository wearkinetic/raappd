package responses

import "github.com/hippoai/goerr"

func ErrNoHandler() error {
	return goerr.NewS(ERR_NO_HANDLER)
}
