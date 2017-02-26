package resource

import "github.com/hippoai/raappd/action"

// AddGet
func (r *Resource) AddGet(a *action.Action) {
	_, exists := r.Gets[a.Name]
	if exists {
		FatalGetActionAlreadyExists(a.Name)
	}

	a.Resource = r.Endpoint
	a.Verb = "GET"
	r.Gets[a.Name] = a

}

// SetPost
func (r *Resource) SetPost(a *action.Action) {
	a.Resource = r.Endpoint
	a.Verb = "POST"
	r.Post = a
}

// SetPut
func (r *Resource) SetPut(a *action.Action) {
	a.Resource = r.Endpoint
	a.Verb = "PUT"
	r.Put = a
}

// SetDelete
func (r *Resource) SetDelete(a *action.Action) {
	a.Resource = r.Endpoint
	a.Verb = "DELETE"
	r.Delete = a
}
