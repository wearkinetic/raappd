package action

import (
	"log"

	"github.com/hippoai/goutil"
)

// Action is a wrapper
// that handles authorization, and payload validation for
// the payload and authorization token sent
type Action struct {
	Verb                       string                      `json:"verb"`
	Resource                   string                      `json:"resource"`
	Name                       string                      `json:"name"`
	Description                string                      `json:"description"`
	ClaimsExtractor            ClaimsExtractor             `json:"-"`
	AuthValidator              AuthValidator               `json:"-"`
	Handler                    Handler                     `json:"-"`
	GetDefaultPayload          GetDefaultPayload           `json:"-"`
	GetParsedPayload           func() (interface{}, error) `json:"-"`
	ExpectedPayloadDescription *PayloadDescription         `json:"expectedPayloadDescription"`
	Extra                      map[string]interface{}      `json:"extra"`
}

// NewAction instanciates
func NewAction(
	name string,
	description string,
	claimsExtractor ClaimsExtractor,
	authValidator AuthValidator,
	handler Handler,
	getDefaultPayload GetDefaultPayload,
) *Action {

	if description == "" {
		FatalNoActionDescription()
	}

	expectedPayloadDescription, err := ExtractExpectedPayload(getDefaultPayload())
	if err != nil {
		log.Fatalf("No payload description | Name %s - Err %s", name, goutil.Pretty(err))
	}

	return &Action{
		Name:                       name,
		Description:                description,
		ClaimsExtractor:            claimsExtractor,
		AuthValidator:              authValidator,
		Handler:                    handler,
		GetDefaultPayload:          getDefaultPayload,
		ExpectedPayloadDescription: expectedPayloadDescription,
		GetParsedPayload: func() (interface{}, error) {
			return nil, ErrParsePayloadNotImplementedYet()
		},
		Extra: map[string]interface{}{},
	}

}
