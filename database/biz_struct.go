package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
)

func FindAllOrgs(creq context.Context, db *sql.DB) (models.BizStructSlice, error) {
	orgs, err := models.BizStructs().All(creq, db)
	return orgs, err
}
