package main

//go:generate sqlboiler -c sqlboiler.toml psql

import (
	"github.com/Shoowa/broker.git/cache"
	data "github.com/Shoowa/broker.git/database"
	server "github.com/Shoowa/broker.git/webserver"
)

func main() {
	db := data.Access()
	red := cache.Setup()
	redjson := cache.SetupRedisJSONClient(red)

	s := server.NewRouter(db, red, redjson)
	s.Run(":8080")
}
