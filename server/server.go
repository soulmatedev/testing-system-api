package server

import (
	"context"
	"github.com/execaus/exloggo"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
	}
	exloggo.Info("server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(postgres *sqlx.DB, ctx context.Context) {
	exloggo.Info("Server shutdown process started")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		exloggo.Error("Error during HTTP server shutdown: " + err.Error())
	} else {
		exloggo.Info("HTTP listener shutdown successfully")
	}

	if postgres != nil {
		if err := postgres.Close(); err != nil {
			exloggo.Error("Error during database connection close: " + err.Error())
		} else {
			exloggo.Info("Business database connection closed successfully")
		}
	} else {
		exloggo.Info("No database connection to close")
	}

	exloggo.Info("Server shutdown process completed successfully")
}
