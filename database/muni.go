package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FindMuniByName(creq context.Context, db *sql.DB, name string, state string) (int, error) {
	muni, err := models.Companies(
		Select(models.MuniColumns.ID),
		models.MuniWhere.Name.EQ(name),
		models.MuniWhere.StateID.EQ(state),
	).One(creq, db)

	return muni.ID, err
}

func FindMuniByCountyID(creq context.Context, db *sql.DB, countyID int) (models.MuniSlice, error) {
	muni, err := models.Munis(
		models.MuniWhere.CountyID.EQ(countyID),
	).All(creq, db)

	return muni, err
}
