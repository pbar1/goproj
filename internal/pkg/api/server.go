package api

import (
	"github.com/dre1080/fiberlog"
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog/log"
)

type Server struct {
	app  *fiber.App
	port int
}

func NewServer(port int) *Server {
	s := &Server{app: nil, port: port}

	app := fiber.New(&fiber.Settings{
		DisableStartupMessage: true,
		StrictRouting:         true,
		// Views:                 html.NewFileSystem(pkger.Dir("/web"), ".html").AddFunc("queryescape", url.QueryEscape),
	})
	app.Use(fiberlog.New(fiberlog.Config{
		Logger: &log.Logger,
	}))

	app.Get("/", s.handleIndex)

	s.app = app
	return s
}

func (s *Server) Start() error {
	log.Info().Msgf("Server listening on 0.0.0.0:%d", s.port)
	return s.app.Listen(s.port)
}
