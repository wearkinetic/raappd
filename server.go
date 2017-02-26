package raappd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hippoai/raappd/resource"
	cors "github.com/itsjamie/gin-cors"
	"github.com/wearkinetic/beasag/responses"
)

type Server struct {
	Engine    *gin.Engine
	Resources map[string]*resource.Resource
}

// NewServer instanciates a server
// With Recovery (can't crash)
// And Logger
// Accepting connections from everywhere (CORS)
func NewServer() *Server {

	// New Gin engine
	// With Recovery (can't crash)
	// And Logger
	// Accepting connections from everywhere
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		RequestHeaders: "Content-Type, Authorization",
	}))

	// Ping
	engine.GET("/ping", func(c *gin.Context) {
		responses.RespondObject(c, map[string]interface{}{
			"message": "pong",
		})
	})

	return &Server{
		Engine:    engine,
		Resources: map[string]*resource.Resource{},
	}

}

// Run serves on a port
func (server *Server) Run(port int) {
	server.Engine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
