package resource

import (
	"fmt"

	"github.com/hippoai/raappd/action"
)

type Doc struct {
	Endpoint        string                 `json:"endpoint"`
	Description     string                 `json:"description"`
	Gets            map[string]*action.Doc `json:"gets"`
	Post            *action.Doc            `json:"post"`
	Put             *action.Doc            `json:"update"`
	Delete          *action.Doc            `json:"delete"`
	Patch           *action.Doc            `json:"patch"`
	ResponseFormats []*ResponseFormat      `json:"responseFormats"`
}

func (r *Resource) MakeDoc() *Doc {

	doc := &Doc{
		Endpoint:        r.Endpoint,
		Description:     r.Description,
		ResponseFormats: r.ResponseFormats,
	}

	// Add the gets documentation
	gets := map[string]*action.Doc{}
	for i, get := range r.Gets {
		gets[fmt.Sprintf("GET.%d", i)] = get.MakeDoc()
	}
	doc.Gets = gets

	// Post
	if r.Post != nil {
		doc.Post = r.Post.MakeDoc()
	}

	// Put
	if r.Put != nil {
		doc.Put = r.Put.MakeDoc()
	}

	// Delete
	if r.Delete != nil {
		doc.Delete = r.Delete.MakeDoc()
	}

	// Patch
	if r.Patch != nil {
		doc.Patch = r.Patch.MakeDoc()
	}

	return doc
}
