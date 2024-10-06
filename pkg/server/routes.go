package server

import (
	"net/http"


	inertia "github.com/romsar/gonertia"
)

// Accept other services as argument
func SetupRoutes(i *inertia.Inertia, c *Container) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(DashboardHandler(i)))

	// Auth
	mux.Handle("GET /login", i.Middleware(LoginGetHandler(i)))
	mux.Handle("POST /login", i.Middleware(LoginPostHandler(i)))
	mux.Handle("DELETE /logout", i.Middleware(LogoutDeleteHandler(i)))

	// Organizations
	mux.Handle("GET /organizations", i.Middleware(OrganizationsGetHandler(i, c)))
	mux.Handle("GET /organizations/create", i.Middleware(OrganizationsCreateGetHandler(i)))
	mux.Handle("POST /organizations", i.Middleware(OrganizationsCreatePostHandler(i, c)))

	// Static files
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	return mux
}

func middlewareSharedProp(i *inertia.Inertia) {
	i.ShareProp("auth", map[string]interface{}{
		"user": map[string]interface{}{
			"first_name": "John",
			"last_name":  "Doe",
			"account": map[string]interface{}{
				"name": "Acme Corporation",
			},
		},
	})
}