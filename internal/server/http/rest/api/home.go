package api

import (
	"net/http"

	"github.com/sluga1naroda/template-web-server/internal/config"

	"github.com/gin-gonic/gin"
)

func GetHome(router *gin.Engine) {
	router.GET("/home", func(ctx *gin.Context) {
		info := config.Get()
		ctx.HTML(http.StatusOK, "home.html", info)
	})
}
