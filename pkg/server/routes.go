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
	oh := newOrganizationsHandler(c, i)
	mux.Handle("GET /organizations", ima.Then(http.HandlerFunc(oh.GetHandler)))
	mux.Handle("GET /organizations/create", ima.Then(http.HandlerFunc(oh.CreateGetHandler)))
	mux.Handle("GET /organizations/{id}/edit", ima.Then(http.HandlerFunc(oh.EditGetHandler)))
	mux.Handle("POST /organizations", ima.Then(http.HandlerFunc(oh.CreatePostHandler)))

	// Static files
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	return mux
}
