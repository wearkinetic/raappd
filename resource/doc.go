package resource

import "github.com/hippoai/raappd/action"

type Doc struct {
	Endpoint    string                 `json:"endpoint"`
	Description string                 `json:"description"`
	Gets        map[string]*action.Doc `json:"gets"`
	Post        *action.Doc            `json:"post"`
	Put         *action.Doc            `json:"update"`
	Delete      *action.Doc            `json:"delete"`
}

func (r *Resource) MakeDoc() *Doc {

	doc := &Doc{
		Endpoint:    r.Endpoint,
		Description: r.Description,
	}

	// Add the gets documentation
	gets := map[string]*action.Doc{}
	for key, get := range r.Gets {
		gets[key] = get.MakeDoc()
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

	return doc
}
