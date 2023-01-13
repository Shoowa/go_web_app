package main

//go:generate sqlboiler -c sqlboiler.toml psql

import (
	"github.com/Shoowa/broker.git/cache"
	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/security"
	server "github.com/Shoowa/broker.git/webserver"
)

func main() {
	awsSess := security.CreateAwsSession()

	db := data.Access()
	red := cache.Setup()
	redjson := cache.SetupRedisJSONClient(red)

	s := server.NewRouter(db, red, redjson, awsSess)
	s.Run(":8080")
}
