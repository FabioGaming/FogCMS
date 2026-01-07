package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"fogcms-server/internal/modules"
	"fogcms-server/internal/modules/hello"
)

// NewRouter creates and configures a new Gin engine with middleware and modules.
func NewRouter() http.Handler {
	router := gin.Default()

	// Apply middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Register all modules
	registerModules(router,
		hello.NewModule(),
	)

	return router
}

// registerModules registers all provided modules with the router.
func registerModules(router *gin.Engine, mods ...modules.Module) {
	for _, mod := range mods {
		mod.RegisterRoutes(router)
	}
}
