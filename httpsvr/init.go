package httpsvr

import (
	"net/http"
	"os"
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

	s.server.Logger.Fatal(s.server.StartServer(svr))
}

func New() IHTTPServer {
	return &HTTPServer{}
}
