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

	return &Server{
		Engine:    engine,
		Resources: map[string]*resource.Resource{},
	}

}

// AddDefaultEndpoints
func (server *Server) AddDoc() {

	// Try to make doc, if something is incomplete it will fail here
	server.MakeDoc()

	// Doc
	server.Engine.GET("/doc", func(c *gin.Context) {
		responses.RespondObject(c, server.MakeDoc())
	})

}

// Run serves on a port
func (server *Server) Run(port int) {
	// Ping
	server.Engine.GET("/ping", func(c *gin.Context) {
		responses.RespondObject(c, map[string]interface{}{
			"message": "pong",
		})
	})

	// AddDoc
	server.AddDoc()

	// Run
	server.Engine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
