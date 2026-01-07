package modules

import "github.com/gin-gonic/gin"

// Module defines the interface for feature modules.
// Each module is responsible for registering its own routes.
type Module interface {
	RegisterRoutes(router *gin.Engine)
}
