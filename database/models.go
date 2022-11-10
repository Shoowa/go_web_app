package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateModel(creq context.Context, db *sql.DB, model models.Model) error {
	err := model.Insert(creq, db, boil.Infer())
	return err
}
