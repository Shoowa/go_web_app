package security

//go:generate go run gen.go

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shoowa/broker.git/models"
	"github.com/aws/aws-sdk-go/aws"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	secMan "github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/caarlos0/env/v6"
	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	null "github.com/volatiletech/null/v8"
)

type config struct {
	Domain     string `env:"COOKIE_DOMAIN,notEmpty" envDefault:"localhost"`
	Auth       string `env:"SECRET_COOKIE_AUTH,notEmpty"`
	Encrypt    string `env:"SECRET_COOKIE_ENCRYPT,notEmpty"`
	AwsProfile string `env:"AWS_PROFILE,notEmpty"`
}

func readConfig() config {
	cfg := config{}
	if env.Parse(&cfg) != nil {
		log.Fatal("Error parsing environ variables for Security configuration.")
	}
	return cfg
}

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

func createCookieStore(configSec config, auth []byte, encrypt []byte) *sessions.CookieStore {
	cookieStore := sessions.NewCookieStore(auth, encrypt)
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.MaxAge = 14400
	cookieStore.Options.SameSite = http.SameSiteStrictMode
	cookieStore.Options.Domain = configSec.Domain
	cookieStore.Options.Secure = true
	return cookieStore
}

func CookieAuthKey() []byte {
	return securecookie.GenerateRandomKey(64)
}

func CookieEncryptionKey() []byte {
	return securecookie.GenerateRandomKey(32)
}

func CreateAwsSession() *awsSession.Session {
	configSec := readConfig()

	options := awsSession.Options{
		Profile:           configSec.AwsProfile,
		SharedConfigState: awsSession.SharedConfigEnable,
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}

	return awsSession.Must(awsSession.NewSessionWithOptions(options))
}

func SvcSecretsManager(sess *awsSession.Session) secMan.SecretsManager {
	return *secMan.New(sess)
}

func readSecret(awsSM *secMan.SecretsManager, secretName string) string {
	name := &secMan.GetSecretValueInput{SecretId: &secretName}
	result, err := awsSM.GetSecretValue(name)

	if err != nil {
		log.Fatal(err.Error())
	}

	return *result.SecretString
}

func SecureCookieStore(awsSM *secMan.SecretsManager) *sessions.CookieStore {
	configSec := readConfig()
	auth := readSecret(awsSM, configSec.Auth)
	encrypt := readSecret(awsSM, configSec.Encrypt)
	fmt.Println(auth)
	fmt.Println([]byte(auth))
	return createCookieStore(configSec, []byte(auth), []byte(encrypt))
}
