package data

import (
	"encoding/json"
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
		CompanyID:  rand.Intn(10),
		Sku:        randomSku,
		Traits:     nullJSON,
		CommonID:   16,
		CategoryID: 17,
		BrandID:    null.IntFrom(rand.Intn(20)),
	}

	return p
}

func RandomProductJSON() []byte {
	p := randomProduct()
	payload, _ := json.Marshal(p)
	return payload
}
