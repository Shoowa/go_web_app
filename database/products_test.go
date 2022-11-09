package data

import (
	"context"
	"testing"

	"github.com/Shoowa/broker.git/models"
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
	db := Access()

	catalog, _ := ReadProducts(context.Background(), db)

	for _, product := range catalog {
		for _, item := range items {
			if product.Company == item.company && product.Sku == item.sku {
				assert.Equal(t, product.Price, item.price)
			}
		}
	}

}

func TestReadProductsFromCompany(t *testing.T) {
	db := Access()

	catalog, _ := ReadProductsFromCompany(context.Background(), db, "Frenchy Laundering")

	for _, product := range catalog {

		// Cotton, 800 thread count, Yellow, Full-Size
		if product.Sku == "C800YEL0F" {
			assert.Equal(t, product.Price, 700)
			break
		}
	}

}

func TestFindProductsByCompanyId(t *testing.T) {
	db := Access()

	products, _ := FindProductsByCompanyId(context.Background(), db, 12)

	for _, product := range products {
		assert.Equal(t, product.CompanyID, 12)
	}
}

func TestCreateProduct(t *testing.T) {
	db := Access()

	newProduct := randomProduct()
	err := CreateProduct(context.Background(), db, newProduct)
	assert.Equal(t, err, nil)

	product, _ := models.Products(
		models.ProductWhere.CompanyID.EQ(newProduct.CompanyID),
		models.ProductWhere.Sku.EQ(newProduct.Sku),
	).One(context.Background(), db)
	defer product.Delete(context.Background(), db)

	assert.Equal(t, product.Sku, newProduct.Sku)

}
