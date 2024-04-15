package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pretender/pkg/logger"
)

type ContentType string

func (c ContentType) String() string {
	return string(c)
}

const (
	JSONContentType ContentType = "application/json"
	HTMLContentType ContentType = "text/html"
	TextContentType ContentType = "text/plain"
)

type Server struct {
	app            *fiber.App
	port           int
	contentType    ContentType
	defaultContent []byte
	logger         *zap.Logger
}

func New(port int, contentType ContentType, defaultContent []byte) *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			AppName:               "pretender",
			DisableStartupMessage: true,
			Prefork:               false,
			ReduceMemoryUsage:     true,
		}),
		port:           port,
		contentType:    contentType,
		defaultContent: defaultContent,
		logger:         logger.Named("server"),
	}
}

// Start the server
func (s *Server) Start() error {
	s.app.All("*", func(c *fiber.Ctx) error {
		logger.L().Info("request", zap.String("method", c.Method()), zap.String("path", c.Path()))
		defer logger.L().Info("response", zap.String("method", c.Method()), zap.String("path", c.Path()), zap.Int("status", 200))
		c.Set("Content-Type", string(s.contentType))
		if s.defaultContent == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Send(s.defaultContent)
	})
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}

// Close the server
func (s *Server) Close() error {
	return s.app.Shutdown()
}
