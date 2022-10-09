package data

import (
	"context"
	"testing"
	"time"
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

	now := time.Now()
	for _, product := range catalog {
		for _, item := range items {
			if product.Company == item.company && product.Sku == item.sku {
				if product.Price != item.price {
					t.Errorf("Appropriate price of %d not found on %q", item.price, now)
				}
			}
		}
	}

}

func TestReadProductsFromCompany(t *testing.T) {
	c := DefineConfig()
	db := c.Access()

	catalog, _ := ReadProductsFromCompany(context.Background(), db, "Frenchy Laundering")

	now := time.Now()
	for _, product := range catalog {

		// Cotton, 800 thread count, Yellow, Full-Size
		if product.Sku == "C800YEL0F" {

			if product.Price != 700 {
				t.Errorf("Appropriate price of %d not found on %q", 700, now)
			}

			break
		}
	}

}
