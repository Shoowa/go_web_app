package security

import (
	"net/http"
	"os"

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
}

type Ticket struct {
	Stub uuid.UUID
	Info SessionInfo
}

func (t *Ticket) Create(person *models.Person, company *models.Company) *Ticket {

	// add UUID
	t.Stub = uuid.New()

	// add employee data
	t.Info.Email = person.Email
	t.Info.ID = person.ID
	t.Info.CompanyID = person.CompanyID
	t.Info.Agent = person.Agent

	// add company data
	t.Info.StateID = company.StateID
	t.Info.CountyID = company.CountyID
	t.Info.MuniID = company.MuniID

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

func SecureCookieStore() *sessions.CookieStore {
	if os.Getenv("ENVIRON") == "dev" {
		return createCookieStore(CookieAuthKey(), CookieEncryptionKey())
	} else {
		auth := []byte(os.Getenv("COOKIE_AUTH"))
		encrypt := []byte(os.Getenv("COOKIE_ENCRYPT"))
		return createCookieStore(auth, encrypt)
	}
}
