package data

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
)

type stuff struct {
	company string
	sku     string
	price   int
}

var items = []stuff{
	stuff{"Frenchy Laundering", "C800YEL0F", 700},
	stuff{"Cheap Cups", "CLAYW01", 1200},
	stuff{"Frenchy Laundering", "C1000YEL0F", 1050},
	stuff{"Paper Boys", "BIC01R20", 1000},
}

func TestReadProducts(t *testing.T) {
	c := DefineConfig()
	db := c.Access()

	catalog, _ := ReadProducts(db)

	for _, product := range catalog {
		for _, item := range items {
			if product.Company == item.company && product.Sku == item.sku {
				assert.Equal(t, product.Price, item.price)
			}
		}
	}

}

func TestReadProductsFromCompany(t *testing.T) {
	c := DefineConfig()
	db := c.Access()

	catalog, _ := ReadProductsFromCompany(context.Background(), db, "Frenchy Laundering")

	for _, product := range catalog {

		// Cotton, 800 thread count, Yellow, Full-Size
		if product.Sku == "C800YEL0F" {
			assert.Equal(t, product.Price, 700)
			break
		}
	}

}
