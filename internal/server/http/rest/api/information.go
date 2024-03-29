package api

import (
	"net/http"

	"github.com/sluga1naroda/template-web-server/internal/config"

	"github.com/gin-gonic/gin"
)

func GetInfo(router *gin.Engine) {
	router.GET("", func(ctx *gin.Context) {
		info := config.Get()
		ctx.HTML(http.StatusOK, "about_me.html", info)
	})
}
