package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/hippoai/raappd/responses"
)

// MakeGetsHandler
// combines the gets handler
func (r *Resource) MakeGetsHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		// Look for the action in URI's GET parameters
		queryName := c.Request.URL.Query().Get(GET_NAME)
		a, exists := r.Gets[queryName]
		if !exists {
			responses.RespondNoHandler(c)
			return
		}

		// Return pre-defined handler
		a.GetHandler()(c)
		return

	}

}
