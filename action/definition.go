package action

// Action is a wrapper
// that handles authorization, and payload validation for
// the payload and authorization token sent
type Action struct {
	Verb               string            `json:"verb"`
	Resource           string            `json:"resource"`
	Name               string            `json:"name"`
	Description        string            `json:"description"`
	PayloadDescription string            `json:"payloadDescription"`
	ClaimsExtractor    ClaimsExtractor   `json:"claimsExtractor"`
	PayloadParser      PayloadParser     `json:"payloadParser"`
	AuthValidator      AuthValidator     `json:"authValidator"`
	Handler            Handler           `json:"handler"`
	GetDefaultPayload  GetDefaultPayload `json:"getDefaultPayload"`
}

// NewAction instanciates
func NewAction(
	name string,
	description string,
	payloadDescription string,
	claimsExtractor ClaimsExtractor,
	payloadParser PayloadParser,
	authValidator AuthValidator,
	handler Handler,
	getDefaultPayload GetDefaultPayload,
) *Action {

	if description == "" {
		FatalNoActionDescription()
	}

	return &Action{
		Name:               name,
		Description:        description,
		PayloadDescription: payloadDescription,
		ClaimsExtractor:    claimsExtractor,
		PayloadParser:      payloadParser,
		AuthValidator:      authValidator,
		Handler:            handler,
		GetDefaultPayload:  getDefaultPayload,
	}

}
