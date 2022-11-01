package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func CreateCompany(creq context.Context, db *sql.DB, company models.Company) error {
	err := company.Insert(creq, db, boil.Infer())
	return err
}

func FindCompanyIDByName(creq context.Context, db *sql.DB, name string, state string) (int, error) {
	company, err := models.Companies(
		Select(models.CompanyColumns.ID),
		models.CompanyWhere.Name.EQ(name),
		models.CompanyWhere.StateID.EQ(state),
	).One(creq, db)

	return company.ID, err
}

func FindCompanyById(creq context.Context, db *sql.DB, id int) (*models.Company, error) {
	company, err := models.FindCompany(creq, db, id)
	return company, err
}
