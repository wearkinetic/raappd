package resource

import "github.com/hippoai/raappd/action"

// AddGet
func (r *Resource) AddGet(a *action.Action) {
	_, exists := r.Gets[a.Name]
	if exists {
		FatalGetActionAlreadyExists(a.Name)
	}

	r.Gets[a.Name] = a

}

// SetPost
func (r *Resource) SetPost(action *action.Action) {
	r.Post = action
}

// SetPut
func (r *Resource) SetPut(action *action.Action) {
	r.Put = action
}

// SetDelete
func (r *Resource) SetDelete(action *action.Action) {
	r.Delete = action
}
