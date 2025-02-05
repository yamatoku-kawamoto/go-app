package web

import (
	"os"

	"github.com/gin-gonic/gin"
)

type (
	Context = gin.Context
)

func SetReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

type Engine struct {
	*gin.Engine
}

func (e *Engine) Run() (err error) {
	host := os.Getenv("HOST")
	if host == "" {
		return e.Engine.Run()
	}
	port := os.Getenv("PORT")
	if port == "" {
		return e.Engine.Run(host + ":8080")
	}
	return e.Engine.Run(host + ":" + port)
}

func New() *Engine {
	return &Engine{
		Engine: gin.Default(),
	}
}

// gin alias
var (
	Dir = gin.Dir
)
