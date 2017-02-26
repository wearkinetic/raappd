package raappd

import "github.com/hippoai/raappd/resource"

// AddResource
func (server *Server) AddResource(r *resource.Resource) {

	_, exists := server.Resources[r.Endpoint]
	if exists {
		FatalResourceAlreadyExists(r.Endpoint)
	}

	server.Resources[r.Endpoint] = r

	// Add the Gets
	if len(r.Gets) > 0 {
		server.Engine.GET(r.Endpoint, r.MakeGetsHandler())
	}

	// Add the Post
	if r.Post != nil {
		server.Engine.POST(r.Endpoint, r.Post.GetHandler())
	}

	// Add the Put
	if r.Put != nil {
		server.Engine.PUT(r.Endpoint, r.Put.GetHandler())
	}

	// Add the Delete
	if r.Delete != nil {
		server.Engine.POST(r.Endpoint, r.Delete.GetHandler())
	}

}
