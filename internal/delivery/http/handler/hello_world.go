package handler

import (
	"TuruGames/internal/types"

	"github.com/gofiber/fiber/v2"
)

type HelloWorldHandler struct {
	Logger *types.Logger
}

func NewHelloWorldHandler(log *types.Logger) *HelloWorldHandler {
	return &HelloWorldHandler{
		Logger: log,
	}
}

func (h *HelloWorldHandler) HelloWorld(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{
		"msg": "hello world",
	})
}
