package web

import (
	"context"
	"fmt"
	"net/http"
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

	server *http.Server
}

func New() *Engine {
	return &Engine{
		Engine: gin.Default(),
	}
}

func (e *Engine) Run() (err error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("%s:%s", host, port)

	e.server = &http.Server{Addr: address, Handler: e.Engine}

	return e.server.ListenAndServe()
}

func (e *Engine) Shutdown(ctx context.Context) (err error) {
	err = e.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

// gin alias
var (
	Dir = gin.Dir
)
