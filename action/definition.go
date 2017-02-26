package action

// Action is a wrapper
// that handles authorization, and payload validation for
// the payload and authorization token sent
type Action struct {
	Verb              string            `json:"verb"`
	Resource          string            `json:"resource"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	ClaimsExtractor   ClaimsExtractor   `json:"claimsExtractor"`
	AuthValidator     AuthValidator     `json:"authValidator"`
	Handler           Handler           `json:"handler"`
	GetDefaultPayload GetDefaultPayload `json:"getDefaultPayload"`
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

	return &Action{
		Name:              name,
		Description:       description,
		ClaimsExtractor:   claimsExtractor,
		AuthValidator:     authValidator,
		Handler:           handler,
		GetDefaultPayload: getDefaultPayload,
	}

}
