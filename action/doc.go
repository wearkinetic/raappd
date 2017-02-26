package action

type Doc struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Parameters  map[string]*ParameterDescription `json:"parameters"`
}

func (a *Action) MakeDoc() *Doc {

	payloadDescription := a.MakePayloadDescription()
	return &Doc{
		Name:        a.Name,
		Description: a.Description,
		Parameters:  payloadDescription,
	}
}
