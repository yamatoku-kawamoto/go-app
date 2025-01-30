//go:build production

package main

import (
	"goapp/internal/repository"
	"goapp/internal/repository/database"
	"goapp/internal/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configuration = Configuration{
		Repository: &repository.Config{
			Database: database.SqliteConfig{
				Path:     "db.sqlite",
				InMemory: false,
			},
		},
	}
}

func initWeb() error {
	web.SetReleaseMode()
	engine = web.New()

	template, err := web.ParseTemplate("views", false, "templates")
	if err != nil {
		return err
	}
	engine.SetHTMLTemplate(template)

	Routing()
	engine.Static("/assets", "views/assets")
	engine.NoRoute(func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.Dir("views", false))
	})

	return nil
}
