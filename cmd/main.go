package main

import (
	"goapp/internal/repository"
	"goapp/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	engine *web.Engine

	configuration Configuration
)

func main() {
	if err := initialize(); err != nil {
		panic(err)
	}
	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func Routing() {
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index", nil)
	})
}

type Configuration struct {
	Repository *repository.Config
}

func (c Configuration) Validate() error {
	return nil
}

func initialize() error {
	if err := initWeb(); err != nil {
		return err
	}
	if err := configuration.Validate(); err != nil {
		return err
	}
	if err := repository.Initialize(configuration.Repository); err != nil {
		return err
	}
	return nil
}
