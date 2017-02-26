package raappd

import (
	"github.com/hippoai/raappd/resource"
)

type Doc struct {
	Resources map[string]*resource.Doc `json:"resources"`
}

func (server *Server) MakeDoc() *Doc {

	doc := &Doc{}

	resources := map[string]*resource.Doc{}
	for key, resource := range server.Resources {
		resources[key] = resource.MakeDoc()
	}
	doc.Resources = resources

	return doc

}
