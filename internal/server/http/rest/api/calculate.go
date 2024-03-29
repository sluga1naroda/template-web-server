package api

import (
	"context"
	"net/http"
	"strconv"

	sllogger "github.com/sluga1naroda/sl-logger"
	"github.com/sluga1naroda/template-web-server/internal/config"
	"github.com/sluga1naroda/template-web-server/internal/model"
	"github.com/sluga1naroda/template-web-server/internal/service"

	"github.com/gin-gonic/gin"
)

func Calculate(router *gin.Engine) {
	router.POST("/calculate", func(ctx *gin.Context) {
		var (
			a, b, c float64
			err     error
		)

		a, err = strconv.ParseFloat(ctx.PostForm("numberA"), 64)
		if err != nil {
			sllogger.Error(context.TODO()).Err(err).Msg("can't define numberA from request")
		}

		b, err = strconv.ParseFloat(ctx.PostForm("numberB"), 64)
		if err != nil {
			sllogger.Error(context.TODO()).Err(err).Msg("can't define numberB from request")
		}

		c, err = strconv.ParseFloat(ctx.PostForm("numberC"), 64)
		if err != nil {
			sllogger.Error(context.TODO()).Err(err).Msg("can't define numberA from request")
		}

		request := model.RequestTask{
			A: a,
			B: b,
			C: c,
		}

		res := service.Calculate(&request)

		res.Name = config.Get().UserInfo.Name
		res.Status = config.Get().UserInfo.Status
		res.Host = config.Get().HTTPServer.Host
		res.Port = strconv.FormatInt(int64(config.Get().HTTPServer.Port), 10)

		ctx.HTML(http.StatusOK, "calc_task_resolved.html", res)
	})
}
