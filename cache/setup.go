package cache

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
	redisJSON "github.com/nitishm/go-rejson/v4"
)

const dot = "."

type config struct {
	Host  string `env:"CACHE_HOST,notEmpty" envDefault:"localhost"`
	Port  string `env:"CACHE_PORT,notEmpty" envDefault:"6379"`
	Pw    string `env:"CACHE_PW,required" envDefault:""`
	DbNum int    `env:"CACHE_NUM,notEmpty" envDefault:"0"`
}

func readConfig() config {
	cfg := config{}
	if env.Parse(&cfg) != nil {
		log.Fatal("Error parsing environ variables for Cache configuration.")
	}
	return cfg
}

func adaptConfig() redis.Options {
	c := readConfig()

	return redis.Options{
		Addr:     c.Host + ":" + c.Port,
		Password: c.Pw,
		DB:       c.DbNum,
	}
}

func Setup() *redis.Client {
	c := adaptConfig()
	return redis.NewClient(&c)
}

func SetupRedisJSONClient(rdb *redis.Client) *redisJSON.Handler {
	rjhandler := redisJSON.NewReJSONHandler()
	rjhandler.SetGoRedisClient(rdb)
	return rjhandler
}
