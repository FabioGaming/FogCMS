package hello

import "github.com/gin-gonic/gin"

// Module represents the hello feature module.
type Module struct {
	handler *Handler
}

// NewModule creates a new hello module with its dependencies.
func NewModule() *Module {
	return &Module{
		handler: NewHandler(),
	}
}

// RegisterRoutes registers all routes for the hello module.
func (m *Module) RegisterRoutes(router *gin.Engine) {
	router.GET("/", m.handler.HelloWorld)
}
