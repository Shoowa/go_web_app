package data

import (
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
	c := PostgresConfig{
		Driver:     "pgx",
		Connection: "user=clerk host=localhost port=60000 dbname=omni",
	}
	db := Open(c)

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
