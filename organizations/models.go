package organizations

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type OrganizationService interface {
	GetOrganization() Organization
	GetOrganizations() []Organization
	CreateOrganization(org Organization) Organization
	DeleteOrganization(id int)
}
