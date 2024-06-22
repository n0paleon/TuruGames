package di

import (
	"TuruGames/internal/delivery/http"
	"TuruGames/internal/delivery/http/handler"
	"TuruGames/internal/types"

	"go.uber.org/dig"
)

func Container(opts *types.Container) {
	container := dig.New()

	// provide default instance
	container.Provide(func() *types.WebFramework { return opts.Server })
	container.Provide(func() *types.Config { return opts.Config })
	container.Provide(func() *types.Logger { return opts.Logger })

	container.Provide(handler.NewHelloWorldHandler)
	container.Provide(http.NewRouter)

	if err := container.Invoke(func(r *http.Router) {
		r.SetupRoutes()
	}); err != nil {
		opts.Logger.Panic("failed to inject dependency: ", err)
	}
}
