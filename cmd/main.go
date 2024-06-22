package cmd

import (
	"TuruGames/infrastructure/config"
	"TuruGames/infrastructure/http"
	"TuruGames/infrastructure/logging"

	"github.com/gofiber/fiber/v2"
)

func StartWebService() {
	config := config.InitConfig()
	logger := logging.NewLogger(config)
	server := http.InitServer(config, logger)

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"msg": "hello world!",
		})
	})
}
