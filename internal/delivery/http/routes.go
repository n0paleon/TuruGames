package http

import "TuruGames/internal/delivery/http/handler"

func (r *Router) ApiRoutesV1() {
	route := r.App.Group("/api/v1")

	hellohandler := r.Handler["hello"].(*handler.HelloWorldHandler)
	mainhandler := r.Handler["main"].(*handler.MainHandler)

	route.Get("/", hellohandler.HelloWorld)
	r.App.Get("/", mainhandler.HelloWorld)
}
