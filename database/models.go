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

func FindModelsByCategoryIDAndBrandID(creq context.Context, db *sql.DB, categoryID int, brandID int) (models.ModelSlice, error) {
	models, err := models.Models(
		models.ModelWhere.CategoryID.EQ(categoryID),
		models.ModelWhere.BrandID.EQ(brandID),
	).All(creq, db)

	if len(models) == 0 {
		return nil, nil
	}

	return models, err
}
