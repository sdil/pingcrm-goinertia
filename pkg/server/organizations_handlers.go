package server

import (
	"encoding/json"
	"net/http"
	"pingcrm/models"
	"pingcrm/organizations"

	"database/sql"

	inertia "github.com/romsar/gonertia"
)

type OrganizationsHandler struct {
	DB *sql.DB
	i  *inertia.Inertia
}

func newOrganizationsHandler(c *Container, i *inertia.Inertia) *OrganizationsHandler {
	h := &OrganizationsHandler{
		DB: c.DB,
		i:  i,
	}
	h.DB = c.DB
	return h
}

func (h *OrganizationsHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	organizations, err := organizations.GetOrganizations(h.DB)
	if err != nil {
		handleServerErr(w, err)
		return
	}

	if organizations == nil {
		organizations = models.OrganizationSlice{}
	}

	err = h.i.Render(w, r, "Organizations/Index", inertia.Props{
		"organizations": map[string]interface{}{
			"data": organizations,
		},
		"filters": map[string]interface{}{
			"search": nil,
		},
	})

	if err != nil {
		handleServerErr(w, err)
	}
}

func (h *OrganizationsHandler) CreateGetHandler(w http.ResponseWriter, r *http.Request) {
	err := h.i.Render(w, r, "Organizations/Create", nil)
	if err != nil {
		handleServerErr(w, err)
		return
	}
}

func (h *OrganizationsHandler) EditGetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	organization, err := organizations.GetOrganization(id, h.DB)

	if err != nil {
		handleServerErr(w, err)
		return
	}

	err = h.i.Render(w, r, "Organizations/Edit", inertia.Props{
		"organization": organization,
	})

	if err != nil {
		handleServerErr(w, err)
		return
	}
}

func (h *OrganizationsHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Organization

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
	h.i.Redirect(w, r, "/organizations")
}
