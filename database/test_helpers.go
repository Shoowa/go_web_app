package data

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/Shoowa/broker.git/models"
	null "github.com/volatiletech/null/v8"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type randomTraits struct {
	Weight int    `json:"weight"`
	Height int    `json:"height"`
	Color  string `'json:"color"'`
}

func randomString(length int, charset string) string {
	rand.Seed(time.Now().UnixMilli())
	b := make([]byte, length)
	max := len(charset)
	for i := range b {
		b[i] = charset[rand.Intn(max)]
	}
	return string(b)
}

func randomProduct() models.Product {
	rand.Seed(time.Now().UnixMilli())
	details := &randomTraits{Weight: rand.Intn(150), Height: rand.Intn(48), Color: "purple"}

	var nullJSON null.JSON
	err := nullJSON.Marshal(details)
	if err != nil {
		panic(err)
	}

	randomSku := randomString(5, charset)

	p := models.Product{
		CompanyID: rand.Intn(10),
		Sku:       randomSku,
		Traits:    nullJSON,
		ModelID:   23,
	}

	return p
}

func RandomProductJSON() []byte {
	p := randomProduct()
	payload, _ := json.Marshal(p)
	return payload
}

func randomCompany() models.Company {
	rand.Seed(time.Now().UnixMilli())
	randomName := randomString(12, charset)

	num := fmt.Sprintf("713-941-%d", rand.Intn(9999))
	addr := fmt.Sprintf("%d Nothing St.", rand.Intn(999))
	p := models.Company{
		Name:      randomName,
		Email:     "help@" + randomName + ".com",
		Telephone: num,
		Street:    addr,
		Postal:    "77075",
		StateID:   "TX",
		CountyID:  1003,
		MuniID:    null.IntFrom(2),
		StructID:  2,
	}

	return p
}

func RandomCompanyJSON() []byte {
	c := randomCompany()
	payload, _ := json.Marshal(c)
	return payload
}

func randomPerson() models.Person {
	rand.Seed(time.Now().UnixMilli())
	randomForename := randomString(6, charset)
	randomSurname := randomString(12, charset)

	p := models.Person{
		CompanyID: rand.Intn(15),
		Forename:  randomForename,
		Surname:   randomSurname,
		Email:     randomForename + "." + randomSurname + "@gmail.com",
		Title:     "purchaser",
		Agent:     true,
	}

	return p
}

func RandomPersonJSON() []byte {
	p := randomPerson()
	payload, _ := json.Marshal(p)
	return payload
}
