package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	sllogger "github.com/sluga1naroda/sl-logger"
)

type Server struct {
	http *http.Server
}

// New the REST API server using the configuration provided.
func New(handler http.Handler, host string, port int) *Server {
	if port < 1 || port > 65535 {
		sllogger.Error(context.TODO()).Msg("server port must be a number between 1 and 65535")
	}

	s := &Server{}
	s.http = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        handler,
		ReadTimeout:    time.Duration(30) * time.Second,
		WriteTimeout:   time.Duration(30) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	return s
}

// Start is a function for start http server.
func (s *Server) Start() {
	if err := s.http.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			sllogger.Info(context.TODO()).Msg("web server shutdown complete")
			return
		}

		sllogger.Fatal(context.TODO()).Msgf("web server closed unexpect: %s", err)
	}
}

// Shutdown is a function for shutdown http server.
func (s *Server) Shutdown(ctx context.Context) error {
	sllogger.Info(context.TODO()).Msg("shutting down web server")

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("web server shutdown failed: %v", err)
	}

	sllogger.Info(context.TODO()).Msg("web server has been closed")

	return nil
}
