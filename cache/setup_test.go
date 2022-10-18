package cache

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-playground/assert/v2"
)

const dot = "."

type Fake struct {
	Name   string `json:"first"`
	Dob    int    `json:"dob"`
	Animal string `json:"animal"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestSetInsertion(t *testing.T) {
	red := Setup()

	data := "{Name: Joshua, Dob: 1987, Animal: ape}"
	err := red.Set(context.Background(), "sess:qwerty123", data, 0).Err()
	check(err)

	val, err := red.Get(context.Background(), "sess:qwerty123").Result()
	check(err)

	assert.Equal(t, val, data)
}

func TestJSONInsertion(t *testing.T) {
	red := Setup()
	redjson := SetupRedisJSONClient(red)

	f := &Fake{Name: "Joshua", Dob: 1987, Animal: "ape"}
	_, err := redjson.JSONSet("123a", dot, f)
	check(err)

	bytes, err := redjson.JSONGet("123a", dot)
	check(err)

	var obj Fake
	json.Unmarshal(bytes.([]byte), &obj)

	assert.Equal(t, obj.Name, f.Name)
}
