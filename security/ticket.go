package security

//go:generate go run gen.go

import (
	"net/http"

	"github.com/Shoowa/broker.git/models"
	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	null "github.com/volatiletech/null/v8"
)

type SessionInfo struct {
	Email     string   `json:"email" binding:"required,email"`
	ID        int      `json:"id" binding:"required"`
	CompanyID int      `json:"company_id" binding:"required"`
	Agent     bool     `json:"agent" binding:"required"`
	StateID   string   `json:"state_id" binding:"required"`
	CountyID  int      `json:"county_id" binding:"required"`
	MuniID    null.Int `json:"muni_id" binding:"required"`
	Forename  string   `json:"forename" binding:"required"`
}

type Ticket struct {
	Stub uuid.UUID
	Info SessionInfo
}

func CreateTicket(person *models.Person, company *models.Company) *Ticket {

	info := SessionInfo{
		Email:     person.Email,
		ID:        person.ID,
		CompanyID: person.CompanyID,
		Agent:     person.Agent,
		StateID:   company.StateID,
		CountyID:  company.CountyID,
		MuniID:    company.MuniID,
		Forename:  person.Forename,
	}

	t := &Ticket{
		Stub: uuid.New(),
		Info: info,
	}

	return t
}

func createCookieStore(auth []byte, encrypt []byte) *sessions.CookieStore {
	cookieStore := sessions.NewCookieStore(auth, encrypt)
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.MaxAge = 14400
	cookieStore.Options.SameSite = http.SameSiteStrictMode
	cookieStore.Options.Domain = "asymblur.com"
	cookieStore.Options.Secure = true
	return cookieStore
}

func CookieAuthKey() []byte {
	return securecookie.GenerateRandomKey(64)
}

func CookieEncryptionKey() []byte {
	return securecookie.GenerateRandomKey(32)
}
