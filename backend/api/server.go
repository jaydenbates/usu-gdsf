package api

import (
	"context"
	"net/http"
	"sync"

	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	wg   *sync.WaitGroup
}

func NewServer(wg *sync.WaitGroup) *Server {
	wg.Add(1)
	return &Server{
		wg: wg,
	}
}

func (s *Server) Start() {
	log.Info("Starting API server")
	s.echo = echo.New()

	s.setupMiddleware()
	s.setupRoutes()

	s.echo.Start(":8080")
}

func (s *Server) Shutdown() {
	log.Info("Shutting down API server")
	s.echo.Shutdown(context.Background())
	log.Info("Done")
	s.wg.Done()
}

func (s *Server) setupMiddleware() {
	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, host=${host}, latency=${latency_human}, error=${error}\n",
	}))

	s.echo.Use(middleware.Gzip())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())

	s.echo.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "/frontend/dist",
		HTML5: true,
	}))
}

func (s *Server) setupRoutes() {
	for _, route := range routes {
		handler := route.handler

		if route.requireAuth || route.requireAdmin {
			handler = auth.RequireAuthorization(route.handler, route.requireAdmin)
		}
		
		switch route.method {
		case http.MethodGet:
			s.echo.GET(route.path, handler)

		case http.MethodPost:
			s.echo.POST(route.path, handler)

		default:
			log.Error("Failed to register unknown method: %v", route.method)
		}
	}
}
