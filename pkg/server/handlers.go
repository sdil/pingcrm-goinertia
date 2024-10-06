package server

import (
	"encoding/json"
	"log"
	"net/http"
	"vuego/organizations"

	inertia "github.com/romsar/gonertia"
)

func DashboardHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewareSharedProp(i)

		err := i.Render(w, r, "Dashboard/Index", nil)
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}

func OrganizationsGetHandler(i *inertia.Inertia, c *Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		organizations := organizations.GetOrganizations(c.DB)

		middlewareSharedProp(i)

		err := i.Render(w, r, "Organizations/Index", inertia.Props{
			"organizations": map[string]interface{}{
				"data": organizations,
			},
			"filters": map[string]interface{}{
				"search": nil,
			},
		})
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}

func OrganizationsCreateGetHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewareSharedProp(i)

		err := i.Render(w, r, "Organizations/Create", inertia.Props{
			"organizations": map[string]interface{}{
				"data": []map[string]interface{}{},
			},
			"filters": map[string]interface{}{
				"search": nil,
			},
		})
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}

func OrganizationsCreatePostHandler(i *inertia.Inertia, c *Container) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req organizations.Organization
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			handleServerErr(w, err)
			return
		}

		organization := organizations.CreateOrganization(req, c.DB)
		i.Redirect(w, r, "/organizations" + string(organization.ID))
	})
}

func handleServerErr(w http.ResponseWriter, err error) {
	log.Printf("http error: %s\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("server error"))
}
