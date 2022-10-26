package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/Shoowa/broker.git/models"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func IsEmailVerified(creq context.Context, db *sql.DB, personID int) (bool, error) {
	clause := models.VerifiedByEmailWhere.SubjectID.EQ(personID)
	sentColumn := models.VerifiedByEmailColumns.Sent

	verity, err := models.VerifiedByEmails(
		clause,
		OrderBy(sentColumn+" DESC"),
	).One(creq, db)

	if err != nil {
		return false, err
	}

	if verity.Response.Valid == true {
		responseTime := verity.Response.Time
		deadline := verity.Sent.Add(5 * time.Minute)
		if verity.Sent.Before(responseTime) && responseTime.Before(deadline) {
			return true, err
		}
	}

	return false, err
}
