package api

import (
	"net/http"

	"github.com/sluga1naroda/template-web-server/internal/config"

	"github.com/gin-gonic/gin"
)

func DrawTask(router *gin.Engine) {
	router.GET("/task", func(ctx *gin.Context) {
		info := config.Get()
		ctx.HTML(http.StatusOK, "calc_task.html", info)
	})
}
