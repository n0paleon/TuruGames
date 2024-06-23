package types

import (
	"TuruGames/infrastructure/logging"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Container struct {
	Config *Config
	Server *WebFramework
	Logger *Logger
}

type Config = viper.Viper
type Logger = logging.SLogger
type WebFramework = fiber.App

type HttpCtx = fiber.Ctx
