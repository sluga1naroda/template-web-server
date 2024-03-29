package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	promhttp "github.com/sluga1naroda/prometheushttp"
	"github.com/sluga1naroda/template-web-server/internal/config"
	"github.com/sluga1naroda/template-web-server/internal/server"
	"github.com/sluga1naroda/template-web-server/internal/server/http/rest"

	"github.com/leandro-lugaresi/hub"
	sllogger "github.com/sluga1naroda/sl-logger"
	cli "github.com/urfave/cli/v2"
)

var StartCommand = &cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts service",
	Action:  startAction,
}

func startAction(ctx *cli.Context) error {
	cfg := config.Get()
	promhttp.SetPrefixForMetrics(cfg.Application.Name)

	handler := rest.NewGin()
	srv := server.New(handler, cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	_ = hub.New()

	// newStore, err := store.New()
	// if err != nil {
	// 	panic(err)
	// }

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go srv.Start()

	<-quit
	sllogger.Get().Info().Msg("shutting down applications")

	cctx, cancel := context.WithTimeout(ctx.Context, time.Duration(30)*time.Second)
	defer cancel()

	if err := srv.Shutdown(cctx); err != nil {
		sllogger.Get().Error().Err(err).Msg("error while server shutdown")
	}

	// newStore.Close(cctx)
	sllogger.Get().Info().Msg("application has been shutted down")

	return nil
}
