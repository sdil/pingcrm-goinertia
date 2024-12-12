package organizations

import (
	"context"
	"database/sql"
	"pingcrm/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetOrganization(id string, db *sql.DB) (*models.Organization, error) {
	ctx := context.Background()
	result, err := models.Organizations(qm.Where("id = ?", id)).One(ctx, db)
	if err != nil {
		return &models.Organization{}, err
	}
	return result, nil
}

func GetOrganizations(db *sql.DB) (models.OrganizationSlice, error) {
	ctx := context.Background()
	result, err := models.Organizations().All(ctx, db)
	if err != nil {
		return models.OrganizationSlice{}, err
	}
	return result, nil
}

func CreateOrganization(org models.Organization, db *sql.DB) (models.Organization, error) {
	ctx := context.Background()
	err := org.Insert(ctx, db, boil.Infer())
	if err != nil {
		return models.Organization{}, err
	}
	return org, nil
}
