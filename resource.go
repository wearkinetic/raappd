package raappd

import (
	"fmt"

	"github.com/hippoai/raappd/resource"
)

// AddResource
func (server *Server) AddResource(version string, r *resource.Resource) {

	versionedEndpoint := fmt.Sprintf("%s/%s", version, r.Endpoint)

	_, exists := server.Resources[versionedEndpoint]
	if exists {
		FatalResourceAlreadyExists(versionedEndpoint)
	}

	server.Resources[versionedEndpoint] = r

	// Add the Gets
	if len(r.Gets) > 0 {
		server.Engine.GET(versionedEndpoint, r.MakeGetHandler())
	}

	// Add the Post
	if r.Post != nil {
		server.Engine.POST(versionedEndpoint, r.MakeNotGetHandler(r.Post))
	}

	// Add the Put
	if r.Put != nil {
		server.Engine.PUT(versionedEndpoint, r.MakeNotGetHandler(r.Put))
	}

	// Add the Delete
	if r.Delete != nil {
		server.Engine.DELETE(versionedEndpoint, r.MakeNotGetHandler(r.Delete))
	}

	// Add the Patch
	if r.Patch != nil {
		server.Engine.PATCH(versionedEndpoint, r.MakeNotGetHandler(r.Patch))
	}

}
