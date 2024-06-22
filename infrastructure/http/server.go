package http

import (
	"TuruGames/infrastructure/logging"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func InitServer(config *viper.Viper, logger *logging.SLogger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:           config.GetString("app.name"),
		CaseSensitive:     true,
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
		Prefork:           config.GetBool("server.prefork"),
		StrictRouting:     true,
		WriteTimeout:      30 * time.Second,
		ErrorHandler:      HttpErrorHandler(config, logger),
	})

	return app
}
