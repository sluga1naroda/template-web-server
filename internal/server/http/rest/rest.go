package rest

import (
	"github.com/sluga1naroda/template-web-server/internal/config"
	"github.com/sluga1naroda/template-web-server/internal/server/http/rest/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	promhttp "github.com/sluga1naroda/prometheushttp"
)

// NewGin is a function for creating routes.
func NewGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	if config.Get().Application.Debug {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	router.LoadHTMLFiles("./internal/resources/about_me.html",
		"./internal/resources/calc_task.html",
		"./internal/resources/calc_task_resolved.html",
		"./internal/resources/home.html",
		"./internal/resources/contacts.html",
	)

	pprof.Register(router)
	router.Use(
		gin.Recovery(),
		cors.Default(),
	)

	router.GET("/metrics", gin.WrapH(promhttp.AddHandler()))

	api.GetInfo(router)
	api.Calculate(router)
	api.GetHome(router)
	api.GetContacts(router)
	api.DrawTask(router)

	router.Use(gin.Logger())

	return router
}
