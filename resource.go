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
		server.Engine.GET(versionedEndpoint, r.MakeGetsHandler())
	}

	// Add the Post
	if r.Post != nil {
		server.Engine.POST(versionedEndpoint, r.Post.GetHandler())
	}

	// Add the Put
	if r.Put != nil {
		server.Engine.PUT(versionedEndpoint, r.Put.GetHandler())
	}

	// Add the Delete
	if r.Delete != nil {
		server.Engine.DELETE(versionedEndpoint, r.Delete.GetHandler())
	}

	// Add the Patch
	if r.Patch != nil {
		server.Engine.PATCH(versionedEndpoint, r.Patch.GetHandler())
	}

}
