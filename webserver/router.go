package webserver

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/Shoowa/broker.git/security"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	redisJSON "github.com/nitishm/go-rejson/v4"
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
	db         *sql.DB
	cache      *redis.Client
	cacheJSON  *redisJSON.Handler
	cookieShop *sessions.CookieStore
}

func NewRouter(db *sql.DB, red *redis.Client, redj *redisJSON.Handler) *gin.Engine {
	p := proxies()

	env := &Env{
		db:         db,
		cache:      red,
		cacheJSON:  redj,
		cookieShop: security.SecureCookieStore(),
	}

	r := gin.Default()
	r.SetTrustedProxies(p)

	r.GET("/health", Health)

	v0 := r.Group("/v0")
	{
		v0.POST(
			"/signup",
			env.MWreadSubmission,
			env.MWregisterCompany,
			env.MWfindCompanyID,
			env.MWregisterPerson,
			env.MWdigestPW,
			env.MWregisterPWDigest,
			env.SignupCompanyPOST,
		)
		v0.POST(
			"/otp",
			env.MWreadEmail,
			env.MWisEmailVerified,
			env.MWdoYouHaveOTP,
			env.MWcreateOTP,
			env.MWshowOTP,
		)
		v0.POST(
			"/login",
			env.MWreadLoginRequest,
			env.MWisEmailVerified,
			env.MWcheckPW,
			env.MWcheckOTP,
			env.MWwriteSessionIntoRedis,
		)
		v0.GET(
			"/welcome",
			env.AuthN,
			env.WelcomeGET,
		)
		v0.DELETE(
			"/logout",
			env.AuthN,
			env.SessionDEL,
		)
		v0.GET("/products", env.ProductsGET)
		v0.GET("/products/:company", env.ProductsFromCompanyGET)
		v0.GET("/productsco/:id", env.ProductsByCompanyIdGET)
		v0.POST("/product", env.CreateProductPOST)
	}

	return r
}
