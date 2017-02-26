package action

func (a *Action) IsPrivate() bool {
	return a.AuthValidator != nil
}
