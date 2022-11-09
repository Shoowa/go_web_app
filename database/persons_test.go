package data

import (
	"context"
	"testing"

	"github.com/Shoowa/broker.git/models"
	"github.com/go-playground/assert/v2"
)

func TestCreatePerson(t *testing.T) {
	db := Access()

	newPerson := randomPerson()
	err := CreatePerson(context.Background(), db, newPerson)
	assert.Equal(t, err, nil)

	person, _ := models.Persons(
		models.PersonWhere.Email.EQ(newPerson.Email),
	).One(context.Background(), db)
	defer person.Delete(context.Background(), db)

	assert.Equal(t, person.Forename, newPerson.Forename)
}
