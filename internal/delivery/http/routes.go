package http

func (r *Router) ApiRoutesV1() {
	route := r.App.Group("/api/v1")

	route.Get("/", r.Handler.HelloWorld.HelloWorld)
}
