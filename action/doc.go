package action

type Doc struct {
	Name                      string                 `json:"name"`
	Description               string                 `json:"description"`
	ExpectedPayloadDesription *PayloadDescription    `json:"parameters"`
	Extra                     map[string]interface{} `json:"extra"`
}

func (a *Action) MakeDoc() *Doc {

	return &Doc{
		Name:                      a.Name,
		Description:               a.Description,
		ExpectedPayloadDesription: a.ExpectedPayloadDescription,
		Extra: a.Extra,
	}
}

func (a *Action) AppendToDoc(key string, value interface{}) {
	a.Extra[key] = value
}
