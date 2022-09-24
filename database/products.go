package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/Shoowa/broker.git/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ProductIDAndPrice struct {
	ID      int    `boil:"products.id" json:"id"`
	Company string `boil:"companies.name" json:"company"`
	Sku     string `boil:"products.sku" json:"sku"`
	Price   int    `boil:"prices.price" json:"price"`
}

type CatalogWithID []*ProductIDAndPrice

func ReadProducts(db *sql.DB) (CatalogWithID, error) {
	ctx := context.Background()

	now := time.Now()
	var catalog CatalogWithID
	err := models.Products(
		Select("DISTINCT ON (companies.name, products.sku) products.id, companies.name, products.sku, prices.price"),
		InnerJoin("companies ON companies.id = products.company_id"),
		InnerJoin("prices ON prices.product_id = products.id"),
		models.PriceWhere.Since.LT(now),
		OrderBy("companies.name, products.sku, prices.since DESC"),
	).Bind(ctx, db, &catalog)
	fail(err)

	return catalog, nil
}
