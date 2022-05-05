package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HelloWorld(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Hello, %s", name)})
}
