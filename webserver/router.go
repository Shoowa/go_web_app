package webserver

import (
	"database/sql"
	"net/http"
	"os"

	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
)

func proxies() []string {
	environ := os.Getenv("ENVIRON")
	var p []string

	if environ == "dev" {
		p = append(p, "127.0.0.1")
	} else {
		p = append(p,
			"10.0.48.0/24",
			"10.0.49.0/24",
			"10.0.50.0/24",
		)
	}

	return p
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

type Env struct {
	db *sql.DB
}

func (e *Env) readProducts(c *gin.Context) {
	catalog, err := data.ReadProducts(e.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "bad"})
	}

	c.JSON(http.StatusOK, catalog)
}

func NewRouter(db *sql.DB) *gin.Engine {
	p := proxies()
	env := &Env{db: db}

	r := gin.Default()
	r.SetTrustedProxies(p)

	r.GET("/health", Health)

	v0 := r.Group("/v0")
	{
		v0.GET("/products", env.readProducts)

	}

	return r
}
