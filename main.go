package main

//go:generate sqlboiler -c sqlboiler.toml psql

import (
	data "github.com/Shoowa/broker.git/database"
	server "github.com/Shoowa/broker.git/webserver"
)

func main() {
	c := data.DefineConfig()
	db := c.Access()

	s := server.NewRouter(db)
	s.Run(":8080")
}
