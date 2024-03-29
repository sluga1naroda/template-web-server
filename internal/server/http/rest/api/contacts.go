package api

import (
	"net/http"

	"github.com/sluga1naroda/template-web-server/internal/config"

	"github.com/gin-gonic/gin"
)

func GetContacts(router *gin.Engine) {
	router.GET("/contacts", func(ctx *gin.Context) {
		info := config.Get()
		ctx.HTML(http.StatusOK, "contacts.html", info)
	})
}
