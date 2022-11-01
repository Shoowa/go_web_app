package data

import (
	"context"
	"database/sql"

	"github.com/Shoowa/broker.git/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateOtpKey(creq context.Context, db *sql.DB, email string, secret string) error {
	record := models.Otp{Email: email, Totpsecret: secret}
	err := record.Insert(creq, db, boil.Infer())
	return err
}

func OtpSecretExists(creq context.Context, db *sql.DB, email string) (bool, error) {
	exist, err := models.Otps(
		models.OtpWhere.Email.EQ(email),
	).Exists(creq, db)

	return exist, err
}

func ReadOTP(creq context.Context, db *sql.DB, email string) (*models.Otp, error) {
	OTPCode, err := models.FindOtp(creq, db, email)
	return OTPCode, err
}
