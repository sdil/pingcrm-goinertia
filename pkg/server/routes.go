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

	// Inertia Middleware
	im := alice.New(i.Middleware)
	// Inertia Auth Middleware
	ima := im.Append(am.sharedPropMiddleware)

	// Dashboard
	mux.Handle("/", im.Then(DashboardHandler(i)))

	// Auth
	mux.Handle("GET /login", im.Then(LoginGetHandler(i)))
	mux.Handle("POST /login", im.Then(LoginPostHandler(i)))
	mux.Handle("DELETE /logout", im.Then(LogoutDeleteHandler(i)))

	// Organizations
	oh := newOrganizationsHandler(c, i)
	mux.Handle("GET /organizations", ima.ThenFunc(oh.GetHandler))
	mux.Handle("GET /organizations/create", ima.ThenFunc(oh.CreateGetHandler))
	mux.Handle("GET /organizations/{id}/edit", ima.ThenFunc(oh.EditGetHandler))
	mux.Handle("POST /organizations", ima.ThenFunc(oh.CreatePostHandler))

	// Static files
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	return mux
}
