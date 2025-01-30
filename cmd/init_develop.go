//go:build develop

package main

import (
	"goapp/internal/repository"
	"goapp/internal/repository/database"
	"goapp/internal/web"
	"net/http"
	"os"
)

func init() {
	configuration = Configuration{
		Repository: &repository.Config{
			Database: database.DebugConfig(),
		},
	}
}

func initWeb() error {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8080")

	engine = web.New()

	// develop only
	{
		proxy := web.NewViteProxy("localhost:5173")
		engine.GET("/", proxy.WebSocketHandShake, func(c *web.Context) {
			c.Redirect(http.StatusMovedPermanently, "/index")
		})
		engine.Any("/@vite/", proxy.ServeProxyData)
		engine.Any("/templates/*path", proxy.ServeProxyData)
		engine.Any("/@solid-refresh", proxy.ServeProxyData)
		engine.Any("/node_modules/", proxy.ServeProxyData)
		engine.NoRoute(proxy.ServeProxyData)

		engine.HTMLRender = web.NewViteHTMLRender("localhost:5173")
	}

	Routing()

	return nil
}
