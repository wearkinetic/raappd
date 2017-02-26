package action

type Doc struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	PayloadDescription string `json:"payloadDescription"`
}

func (a *Action) MakeDoc() *Doc {
	return &Doc{
		Name:               a.Name,
		Description:        a.Description,
		PayloadDescription: a.PayloadDescription,
	}
}
