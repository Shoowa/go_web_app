package data

// go:generate_data sqlboiler psql

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Error handling.
func fail(e error) {
	if e != nil {
		panic(e)
	}
}

type PostgresConfig struct {
	Driver     string
	Connection string
}

func DefineConfig() PostgresConfig {
	dbstring := os.Getenv("DBSTRING")
	return PostgresConfig{Driver: "pgx", Connection: dbstring}
}

func (pc PostgresConfig) Access() *sql.DB {
	db, err := sql.Open(pc.Driver, pc.Connection)
	fail(err)
	err = db.Ping()
	fail(err)
	return db
}
