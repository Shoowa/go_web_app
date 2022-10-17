package main

//go:generate sqlboiler -c sqlboiler.toml psql

import (
	"github.com/Shoowa/broker.git/cache"
	data "github.com/Shoowa/broker.git/database"
	server "github.com/Shoowa/broker.git/webserver"
)

func main() {
	c := data.DefineConfig()
	db := c.Access()
	red := cache.Setup()

	s := server.NewRouter(db, red)
	s.Run(":8080")
}
