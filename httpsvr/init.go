package httpsvr

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "1323"
	}

	svr := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := s.server.StartServer(svr); err != nil {
			s.server.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.server.Logger.Fatal(err)
	}
}

// New ...
func New() IHTTPServer {
	return &HTTPServer{}
}
