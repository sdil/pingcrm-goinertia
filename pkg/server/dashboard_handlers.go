package server

import (
	"net/http"
	inertia "github.com/romsar/gonertia"

)

func DashboardHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Dashboard/Index", nil)
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}