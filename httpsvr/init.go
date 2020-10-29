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
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// IHTTPServer is an interface for http server
type IHTTPServer interface {
	Initialize() IHTTPServer
	Server() *echo.Echo
	Start(debug bool)
}

// HTTPServer is an implementation for IHTTPServer
type HTTPServer struct {
	server *echo.Echo
}

// Initialize initiates new http server
func (s *HTTPServer) Initialize() IHTTPServer {
	s.server = echo.New()
	s.server.HideBanner = true
	s.server.Use(middleware.Logger())
	s.server.Use(middleware.Recover())

	s.initiateHealthCheckEndpoint()
	s.initiatePrometheusMetricHandler()

	return s
}

func (s *HTTPServer) initiateHealthCheckEndpoint() {
	s.server.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}

func (s *HTTPServer) initiatePrometheusMetricHandler() {
	s.server.GET("/metrics", func(c echo.Context) error {
		return nil
	}, echo.WrapMiddleware(func(handler http.Handler) http.Handler {
		return promhttp.Handler()
	}))
}

// Server returns echo.Echo object server
func (s *HTTPServer) Server() *echo.Echo {
	return s.server
}

// Start starts http server
func (s *HTTPServer) Start(debug bool) {
	s.server.Debug = debug

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

// New initiates new http server object
func New() IHTTPServer {
	return &HTTPServer{}
}
