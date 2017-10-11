package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwri/go-workflow/api/v1"
	"github.com/kwri/go-workflow/modules/setting"
)

var (
	app = gin.Default()
)

func RegisterRoute() {
	app.GET("/", home)
	v1route := app.Group("/api/v1")
	v1.RegisterRoute(v1route)
	app.Run(setting.Config.ApiPort)
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "kw workflow api",
		"version": "v1.0.0",
	})
}
