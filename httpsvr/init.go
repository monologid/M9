package httpsvr

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/monologid/m9/config"
)

// IHTTPServer ...
type IHTTPServer interface {
	Initialize() IHTTPServer
	Server() *echo.Echo
	Start()
}

// HTTPServer ...
type HTTPServer struct {
	server *echo.Echo
}

// Initialize ...
func (s *HTTPServer) Initialize() IHTTPServer {
	s.server = echo.New()
	s.server.Use(middleware.Logger())
	s.server.Use(middleware.Recover())

	return s
}

// Server returns echo.Echo object server
func (s *HTTPServer) Server() *echo.Echo {
	return s.server
}

// Start ...
func (s *HTTPServer) Start() {
	host := config.C.Application.Host
	port := config.C.Application.Port

	svr := &http.Server{
		Addr:         host + ":" + port,
		ReadTimeout:  config.C.Application.ReadTimeout * time.Second,
		WriteTimeout: config.C.Application.WriteTimeout * time.Second,
	}

	go func() {
		if err := s.server.StartServer(svr); err != nil {
			s.server.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), config.C.Application.GracefulShutdownTimeout*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.server.Logger.Fatal(err)
	}
}

// New ...
func New() IHTTPServer {
	return &HTTPServer{}
}
