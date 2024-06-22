package http

import (
	"TuruGames/infrastructure/logging"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func HttpErrorHandler(config *viper.Viper, logger *logging.SLogger) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		logger.Error("app error: %+v", err.Error())

		if !config.GetBool("app.development") {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"error": "error while processing your request",
			})
		}

		return ctx.Status(code).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
}
