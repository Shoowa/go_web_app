package cache

import (
	"os"

	"github.com/go-redis/redis/v8"
	redisJSON "github.com/nitishm/go-rejson/v4"
)

const dot = "."

func config() redis.Options {
	if os.Getenv("ENVIRON") == "dev" {
		return redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}
	}

	host := os.Getenv("CACHEHOST")
	port := os.Getenv("CACHEPORT")
	pw := os.Getenv("CACHEPW")
	db := 0

	return redis.Options{
		Addr:     host + ":" + port,
		Password: pw,
		DB:       db,
	}
}

func Setup() *redis.Client {
	c := config()
	return redis.NewClient(&c)
}

func SetupRedisJSONClient(rdb *redis.Client) *redisJSON.Handler {
	rjhandler := redisJSON.NewReJSONHandler()
	rjhandler.SetGoRedisClient(rdb)
	return rjhandler
}
