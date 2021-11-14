package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"paper.com/paper/config"
	"paper.com/paper/middleware"
	"paper.com/paper/route"
)

type Server struct {
	e      *echo.Echo
	config *config.Config
}

func SetupServer(c *config.Config) *Server {
	server := &Server{}

	// Initialize Echo
	e := echo.New()

	// Set Middleware
	e.Use(middleware.SetCommonHeaders)

	// Set Route
	route.InitRoute(e)

	// Set Server parameters
	server.e = e
	server.config = c

	return server
}

func (s *Server) Start() {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", s.config.Server.Port)))
}
