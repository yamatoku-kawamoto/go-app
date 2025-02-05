package web_test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct{}

func (h *ExampleHandler) Page(c *gin.Context) {
	c.HTML(http.StatusOK, "example/template", nil)
}
