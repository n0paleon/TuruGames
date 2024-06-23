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
	container.Provide(func() *types.Logger { return opts.Logger })
	container.Provide(func() *types.WebFramework { return opts.Server })
	container.Provide(func() *types.Config { return opts.Config })

	container.Provide(handler.NewHelloWorldHandler)
	container.Provide(handler.NewMainHandler)

	// inject handler to router
	container.Provide(func(
		h1 *handler.HelloWorldHandler,
		h2 *handler.MainHandler,
	) http.Handler {
		return http.Handler{
			"hello": h1,
			"main":  h2,
		}
	})

	container.Provide(http.NewRouter)

	if err := container.Invoke(func(r *http.Router) {
		r.SetupRoutes()
	}); err != nil {
		opts.Logger.Panic("failed to inject dependency: ", err)
	}
}
