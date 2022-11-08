package data

// go:generate_data sqlboiler psql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Error handling.
func fail(e error) {
	if e != nil {
		panic(e)
	}
}

type config struct {
	User   string `env:"BROKER_DBUSER,notEmpty"`
	Host   string `env:"BROKER_DBHOST,notEmpty"`
	Port   string `env:"BROKER_DBPORT,notEmpty"`
	DbName string `env:"BROKER_DBNAME,notEmpty"`
}

func readConfig() config {
	cfg := config{}
	if env.Parse(&cfg) != nil {
		log.Fatal("Error parsing environ variables for Postgres configuration.")
	}
	return cfg
}

func (c config) dbString() string {
	return fmt.Sprintf(
		"user=%s host=%s port=%s dbname=%s",
		c.User, c.Host, c.Port, c.DbName,
	)
}

func Access() *sql.DB {
	dbConnection := readConfig().dbString()
	db, err := sql.Open("pgx", dbConnection)
	fail(err)
	err = db.Ping()
	fail(err)
	return db
}
