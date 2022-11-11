//go:build ignore

package main

import (
	"github.com/Shoowa/broker.git/security"
	"log"
	"os"
	"text/template"
	"time"
)

type genData struct {
	Timestamp  time.Time
	Auth       []byte
	Encryption []byte
}

func main() {
	f, err := os.Create("cookie_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	packageTemplate.Execute(f, &genData{
		Timestamp:  time.Now(),
		Auth:       security.CookieAuthKey(),
		Encryption: security.CookieEncryptionKey(),
	})
}

var packageTemplate = template.Must(template.New("").Parse(`// DO NOT EDIT generated code
// WHEN: {{ .Timestamp }}
package security

import (
	"os"
	"github.com/gorilla/sessions"
)

type CookieKeys struct {
	Auth		[]byte
	Encryption	[]byte
}

var CookieKeysConfig = CookieKeys{
	Auth: []byte({{ printf "%q" .Auth }}),
	Encryption: []byte({{ printf "%q" .Encryption }}),
}

func SecureCookieStore() *sessions.CookieStore {
	if os.Getenv("ENVIRON") == "dev" {
		return createCookieStore(CookieKeysConfig.Auth, CookieKeysConfig.Encryption)
	} else {
		auth := []byte(os.Getenv("COOKIE_AUTH"))
		encrypt := []byte(os.Getenv("COOKIE_ENCRYPT"))
		return createCookieStore(auth, encrypt)
	}
}
`))
