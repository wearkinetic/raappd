// package resource
// wraps a REST endpoint for a given resource
// with access to Gets methods, indexed by specific action
// and Post, Put, Delete for the given resource
package resource

import (
	"github.com/hippoai/raappd/action"
)

// Resource wraps all the get, post, update and delete "verb"s for
// for a given resource. It also enforces a description.
type Resource struct {
	Endpoint    string                    `json:"endpoint"`
	Description string                    `json:"description"`
	Gets        map[string]*action.Action `json:"gets"`
	Post        *action.Action            `json:"post"`
	Put         *action.Action            `json:"update"`
	Delete      *action.Action            `json:"delete"`
	Patch       *action.Action            `json:"patch"`
}

// NewResource instanciates an Endpoint
// It forces to have a resource name and a description
func NewResource(endpoint, description string) *Resource {
	if endpoint == "" {
		FatalNoResourceEndpoint()
	}
	if description == "" {
		FatalNoResourceDescription()
	}

	return &Resource{
		Endpoint:    endpoint,
		Description: description,
		Gets:        map[string]*action.Action{},
		Post:        nil,
		Put:         nil,
		Delete:      nil,
		Patch:       nil,
	}
}
