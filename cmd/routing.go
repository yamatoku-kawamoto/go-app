package main

import (
	"goapp/internal/web"
	"net/http"
)

func Routing() {
	engine.GET("/index", func(c *web.Context) {
		c.HTML(http.StatusOK, "index/index", nil)
	})
}
