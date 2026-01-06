package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler contains all HTTP handlers for the hello module.
type Handler struct{}

// NewHandler creates a new Handler instance.
func NewHandler() *Handler {
	return &Handler{}
}

// HelloWorld handles GET / and returns a hello world message.
func (h *Handler) HelloWorld(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
