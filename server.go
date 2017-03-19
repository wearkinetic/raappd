package raappd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hippoai/raappd/resource"
	"github.com/hippoai/raappd/responses"
	cors "github.com/itsjamie/gin-cors"
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
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Content-Type, Authorization, Content-Range",
	}))

	return &Server{
		Engine:    engine,
		Resources: map[string]*resource.Resource{},
	}

}

// Run serves on a port
func (server *Server) Run(port int) {
	// Ping
	server.Engine.GET("/ping", func(c *gin.Context) {
		responses.RespondObject(c, map[string]interface{}{
			"message": "pong",
		})
	})

	// MakeDoc
	server.MakeDoc()

	// Run
	server.Engine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
