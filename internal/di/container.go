package di

import (
	"TuruGames/internal/types"

	"go.uber.org/dig"
)

func Container(opts *types.Container) {
	container := dig.New()

	// provide default instance
	container.Provide(func() *types.WebFramework { return opts.Server })
	container.Provide(func() *types.Config { return opts.Config })
	container.Provide(func() *types.Logger { return opts.Logger })
}
