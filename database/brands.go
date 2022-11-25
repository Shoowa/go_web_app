package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func FindAllBrands(creq context.Context, db *sql.DB) (models.BrandSlice, error) {
	brands, err := models.Brands().All(creq, db)
	return brands, err
}

type BrandsFromModelID struct {
	ID   int    `boil:"models.brand_id" json:"id"`
	Name string `boil:"brands.name" json:"name"`
}

type EligibleBrands []*BrandsFromModelID

func FindBrandsRelatedToModelID(creq context.Context, db *sql.DB, modelID int) (EligibleBrands, error) {
	brandID := models.ModelColumns.BrandID
	brandName := models.BrandColumns.Name
	desiredColumns := Select(brandID, brandName)

	var brands EligibleBrands
	err := models.Models(
		desiredColumns,
		InnerJoin("brands ON brands.id = models.brand_id"),
		models.ModelWhere.CategoryID.EQ(modelID),
		GroupBy(brandID),
		GroupBy(brandName),
	).Bind(creq, db, &brands)

	if len(brands) == 0 {
		return nil, nil
	}

	return brands, err
}
