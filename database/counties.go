package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FindCountyByName(creq context.Context, db *sql.DB, name string, state string) (int, error) {
	county, err := models.Counties(
		Select(models.CountyColumns.ID),
		models.CountyWhere.Name.EQ(name),
		models.CountyWhere.StateID.EQ(state),
	).One(creq, db)

	return county.ID, err
}

func FindCountiesByStateID(creq context.Context, db *sql.DB, stateID string) (models.CountySlice, error) {
	counties, err := models.Counties(
		models.CountyWhere.StateID.EQ(stateID),
	).All(creq, db)

	return counties, err
}
