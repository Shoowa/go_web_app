package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
)

func FindAllBrands(creq context.Context, db *sql.DB) (models.BrandSlice, error) {
	brands, err := models.Brands().All(creq, db)
	return brands, err
}
