package server

import (
	"net/http"

	inertia "github.com/romsar/gonertia"
)

func LoginGetHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Auth/Login", nil)
		if err != nil {
			return
		}
	})
}

func LoginPostHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i.Redirect(w, r, "/")
	})
}

func LogoutDeleteHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i.Redirect(w, r, "/login")
	})
}
