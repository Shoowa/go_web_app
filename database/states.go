package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
)

func FindAllStates(creq context.Context, db *sql.DB) (models.StateSlice, error) {
	states, err := models.States().All(creq, db)
	return states, err
}
