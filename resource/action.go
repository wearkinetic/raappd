package resource

import (
	"sort"

	"github.com/hippoai/raappd/action"
)

// AddGet
func (r *Resource) AddGet(a *action.Action) {

	a.Resource = r.Endpoint
	a.Verb = "GET"
	r.Gets = append(r.Gets, a)

	// Sort - the list of gets - the most complex payload description first
	sort.Sort(r.Gets)

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

// SetPatch
func (r *Resource) SetPatch(a *action.Action) {
	a.Resource = r.Endpoint
	a.Verb = "PATCH"
	r.Patch = a
}
