package cache

import (
	"encoding/json"

	redisJSON "github.com/nitishm/go-rejson/v4"
)

const (
	COMPANY_ID = dot + "company_id"
)

func ReadSessionEmail(rj *redisJSON.Handler, stub string) (*string, error) {
	b, err := rj.JSONGet(stub, dot+"email")
	if err != nil {
		return nil, err
	}

	var info string
	json.Unmarshal(b.([]byte), &info)
	return &info, err
}

func ReadSessionForename(rj *redisJSON.Handler, stub string) (*string, error) {
	b, err := rj.JSONGet(stub, dot+"forename")
	if err != nil {
		return nil, err
	}

	var info string
	json.Unmarshal(b.([]byte), &info)
	return &info, err
}

func ReadSessionCompany(rj *redisJSON.Handler, stub string) (*int, error) {
	b, err := rj.JSONGet(stub, COMPANY_ID)
	if err != nil {
		return nil, err
	}

	var info int
	json.Unmarshal(b.([]byte), &info)
	return &info, err
}
