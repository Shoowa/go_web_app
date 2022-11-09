package data

import (
	"context"
	"testing"

	"github.com/Shoowa/broker.git/models"
	"github.com/go-playground/assert/v2"
)

func TestCreateCompany(t *testing.T) {
	db := Access()

	newCompany := randomCompany()
	err := CreateCompany(context.Background(), db, newCompany)
	assert.Equal(t, err, nil)

	company, _ := models.Companies(
		models.CompanyWhere.Name.EQ(newCompany.Name),
	).One(context.Background(), db)
	defer company.Delete(context.Background(), db)

	assert.Equal(t, company.Name, newCompany.Name)
}
