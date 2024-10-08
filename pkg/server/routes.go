package server

import (
	"net/http"
	inertia "pingcrm/pkg/inertia"
)

// Accept other services as argument
func SetupRoutes(c *Container) *http.ServeMux {
	i := inertia.InitInertia()

	mux := http.NewServeMux()

	// Dashboard
	mux.Handle("/", i.Middleware(sharedPropMiddleware(DashboardHandler(i), i)))

	// Auth
	mux.Handle("GET /login", i.Middleware(LoginGetHandler(i)))
	mux.Handle("POST /login", i.Middleware(LoginPostHandler(i)))
	mux.Handle("DELETE /logout", i.Middleware(LogoutDeleteHandler(i)))

	// Organizations
	oh := new(OrganizationsHandler)
	oh.Init(c)
	mux.Handle("GET /organizations", i.Middleware(sharedPropMiddleware(oh.GetHandler(i), i)))
	mux.Handle("GET /organizations/create", i.Middleware(sharedPropMiddleware(oh.CreateGetHandler(i), i)))
	mux.Handle("GET /organizations/{id}/edit", i.Middleware(sharedPropMiddleware(oh.EditGetHandler(i),i)))
	mux.Handle("POST /organizations", i.Middleware(sharedPropMiddleware(oh.CreatePostHandler(i),i)))

	// Static files
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	return mux
}
