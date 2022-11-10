package webserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Shoowa/broker.git/security"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	redisJSON "github.com/nitishm/go-rejson/v4"
)

type config struct {
	Proxies []string `env:"APP_PROXIES,required" envSeparator:","  envDefault:"127.0.0.1"`
}

func readConfig() config {
	cfg := config{}
	if env.Parse(&cfg) != nil {
		log.Fatal("Error parsing environ variables for router configuration.")
	}
	return cfg
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
	cfg := readConfig()

	env := &Env{
		db:         db,
		cache:      red,
		cacheJSON:  redj,
		cookieShop: security.SecureCookieStore(),
	}

	r := gin.Default()
	r.SetTrustedProxies(cfg.Proxies)

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

		authn := v0.Group("/au", env.AuthN)

		authn.GET("/welcome", env.WelcomeGET)
		authn.DELETE("/logout", env.SessionDEL)
		authn.GET("/products", env.NotMyProductsGET)
		authn.GET("/products/:company", env.ProductsFromCompanyGET)
		authn.GET("/productsco/:id", env.ProductsByCompanyIdGET)
		authn.POST("/product", env.CreateProductPOST)
		authn.GET("/myproducts", env.MyProductsGET)
		authn.GET("/models/:code/products", env.FindActiveProductsByModelCodeGET)
		authn.POST("/model", env.CreateModelPOST)
	}

	return r
}
