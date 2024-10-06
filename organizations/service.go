package organizations

import (
	"gorm.io/gorm"
)

func GetOrganization(id string, db *gorm.DB) Organization {
	return Organization{}
}

func GetOrganizations(db *gorm.DB) []Organization {
	var organizations []Organization

	result := db.Find(&organizations)
	if result.Error != nil {
		panic(result.Error)
	}
	return organizations
}

func CreateOrganization(org Organization, db *gorm.DB) Organization {
	result := db.Create(&org)
	if result.Error != nil {
		panic(result.Error)
	}
	return org
}