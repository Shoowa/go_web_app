package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateDigest(creq context.Context, db *sql.DB, digest models.Digest) error {
	err := digest.Insert(creq, db, boil.Infer())
	return err
}

func ReadDigest(creq context.Context, db *sql.DB, email string) (*models.Digest, error) {
	pwDigest, err := models.FindDigest(creq, db, email)
	return pwDigest, err
}
