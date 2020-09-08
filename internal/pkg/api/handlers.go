package api

import "github.com/gofiber/fiber"

func (s *Server) handleIndex(c *fiber.Ctx) {
	c.Send("hello friend")
}
