package http

import (
	"TuruGames/internal/delivery/http/handler"
	"TuruGames/internal/types"
)

type Handler struct {
	HelloWorld *handler.HelloWorldHandler
}

type Router struct {
	App     *types.WebFramework
	Handler *Handler
}

func NewRouter(
	app *types.WebFramework,
	HelloWorldHandler *handler.HelloWorldHandler,
) *Router {
	return &Router{
		App: app,
		Handler: &Handler{
			HelloWorld: HelloWorldHandler,
		},
	}
}

func (r *Router) SetupRoutes() {
	r.ApiRoutesV1()
}
