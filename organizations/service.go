package organizations

import (
	"gorm.io/gorm"
)

func GetOrganization(id string, db *gorm.DB) (Organization, error) {
	var organization Organization
	result := db.Take(&organization, id)
	if result.Error != nil {
		return Organization{}, result.Error
	}
	return organization, nil
}

func GetOrganizations(db *gorm.DB) ([]Organization, error) {
	var organizations []Organization

	result := db.Find(&organizations)
	if result.Error != nil {
		return nil, result.Error
	}
	return organizations, nil
}

func CreateOrganization(org Organization, db *gorm.DB) (Organization, error) {
	result := db.Create(&org)
	if result.Error != nil {
		return Organization{}, result.Error
	}
	return org, nil
}
