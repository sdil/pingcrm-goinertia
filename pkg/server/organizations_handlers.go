package server

import (
	"encoding/json"
	"net/http"
	"pingcrm/organizations"

	"gorm.io/gorm"
	inertia "github.com/romsar/gonertia"
)

type OrganizationsHandler struct {
	DB *gorm.DB
}

func (h *OrganizationsHandler) Init(c *Container) {
	h.DB = c.DB
	println("OrganizationsHandler initialized")
}

func DashboardHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Dashboard/Index", nil)
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}

func (h *OrganizationsHandler) GetHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		organizations, err := organizations.GetOrganizations(h.DB)
		if err != nil {
			handleServerErr(w, err)
			return
		}
		err = i.Render(w, r, "Organizations/Index", inertia.Props{
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

func (h *OrganizationsHandler) CreateGetHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Organizations/Create", nil)
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}

func (h *OrganizationsHandler) EditGetHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		organization, err := organizations.GetOrganization(id, h.DB)
		if err != nil {
			handleServerErr(w, err)
			return
		}
		err = i.Render(w, r, "Organizations/Edit", inertia.Props{
			"organization": organization,
		})
		if err != nil {
			handleServerErr(w, err)
			return
		}
	})
}

func (h *OrganizationsHandler) CreatePostHandler(i *inertia.Inertia) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req organizations.Organization
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			handleServerErr(w, err)
			return
		}

		_, err = organizations.CreateOrganization(req, h.DB)
		if err != nil {
			handleServerErr(w, err)
			return
		}
		i.Redirect(w, r, "/organizations")
	})
}

