package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
)

func FindAllCategories(creq context.Context, db *sql.DB) (models.CategorySlice, error) {
	categories, err := models.Categories().All(creq, db)
	return categories, err
}
