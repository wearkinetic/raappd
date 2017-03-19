package action

type Doc struct {
	Name                      string              `json:"name"`
	Description               string              `json:"description"`
	ExpectedPayloadDesription *PayloadDescription `json:"parameters"`
}

func (a *Action) MakeDoc() *Doc {

	return &Doc{
		Name:                      a.Name,
		Description:               a.Description,
		ExpectedPayloadDesription: a.ExpectedPayloadDescription,
	}
}
