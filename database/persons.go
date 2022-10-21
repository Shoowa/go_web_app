package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreatePerson(creq context.Context, db *sql.DB, person models.Person) error {
	err := person.Insert(creq, db, boil.Infer())
	return err
}
