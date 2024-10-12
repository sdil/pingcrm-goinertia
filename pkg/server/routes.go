package server

import (
	"net/http"
	"github.com/justinas/alice"

	inertia "pingcrm/pkg/inertia"
)

// Accept other services as argument
func SetupRoutes(c *Container) *http.ServeMux {
	i := inertia.InitInertia()
	am := NewAuthMiddleware(i)

	mux := http.NewServeMux()

	im := alice.New(i.Middleware)
	ima := alice.New(i.Middleware, am.sharedPropMiddleware)

	// Dashboard
	mux.Handle("/", im.Then(DashboardHandler(i)))

	// Auth
	mux.Handle("GET /login", im.Then(LoginGetHandler(i)))
	mux.Handle("POST /login", im.Then(LoginPostHandler(i)))
	mux.Handle("DELETE /logout", im.Then(LogoutDeleteHandler(i)))

	// Organizations
	oh := newOrganizationsHandler(c)
	mux.Handle("GET /organizations", ima.Then(oh.GetHandler(i)))
	mux.Handle("GET /organizations/create", ima.Then(oh.CreateGetHandler(i)))
	mux.Handle("GET /organizations/{id}/edit", ima.Then(oh.EditGetHandler(i)))
	mux.Handle("POST /organizations", ima.Then(oh.CreatePostHandler(i)))

	// Static files
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	return mux
}
