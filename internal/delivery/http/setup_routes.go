package http

import (
	"TuruGames/internal/types"
)

type Handler = map[string]interface{}

type Router struct {
	App     *types.WebFramework
	Handler Handler
}

func NewRouter(
	app *types.WebFramework,
	handlers Handler,
) *Router {
	return &Router{
		App:     app,
		Handler: handlers,
	}
}

func (r *Router) SetupRoutes() {
	r.ApiRoutesV1()
}
