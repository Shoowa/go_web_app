package cache

import (
	"encoding/json"

	redisJSON "github.com/nitishm/go-rejson/v4"
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
