package handler

import (
	"TuruGames/internal/types"

	"github.com/gofiber/fiber/v2"
)

type MainHandler struct {
	Logger *types.Logger
}

func NewMainHandler(log *types.Logger) *MainHandler {
	return &MainHandler{
		Logger: log,
	}
}

func (h *MainHandler) HelloWorld(c *types.HttpCtx) error {
	return c.JSON(&fiber.Map{
		"message": "hello world",
		"status":  "success",
	})
}
