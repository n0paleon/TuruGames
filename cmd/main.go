package cmd

import (
	"TuruGames/infrastructure/config"
	"TuruGames/infrastructure/http"
	"TuruGames/infrastructure/logging"
	"TuruGames/internal/di"
	"TuruGames/internal/types"
	"fmt"
)

func StartWebService() {
	config := config.InitConfig()
	logger := logging.NewLogger(config)
	defer logger.Sync()
	server := http.InitServer(config, logger)

	di.Container(&types.Container{
		Config: config,
		Logger: logger,
		Server: server,
	})

	logger.Info("server started!")

	web_host := config.GetString("server.host")
	web_port := config.GetInt("server.port")
	server.Listen(fmt.Sprintf("%s:%d", web_host, web_port))
}
