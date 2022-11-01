package webserver

import (
	"net/http"
	"time"

	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/security"
	"github.com/gin-gonic/gin"
)

type RequestPostLogin struct {
	Email string `json:"email" binding:"required,email"`
	requestCreatePW
	Code string `json:"code" binding:"required"`
}

func (e *Env) MWreadLoginRequest(c *gin.Context) {
	var login RequestPostLogin
	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Set("email", login.Email)
	c.Set("password", login.PW)
	c.Set("code", login.Code)
	c.Next()
}

func (e *Env) MWcheckPW(c *gin.Context) {
	// Use e-mail address to find the digest.
	email := c.GetString("email")
	digest, errDigest := data.ReadDigest(c.Request.Context(), e.db, email)

	// Abort request when digest can't be read from DB.
	if errDigest != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errDigest.Error()})
		return
	}

	// Compare submitted PW to digest.
	pw := c.GetString("password")
	valid := security.CheckHash(pw, digest.Digest)

	// Abort request when password seems invalid.
	if valid == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	c.Next()
}

func (e *Env) MWcheckOTP(c *gin.Context) {
	// Use e-mail address to find OTP Secret.
	email := c.GetString("email")
	OTP, errOTP := data.ReadOTP(c.Request.Context(), e.db, email)

	// Abort request when OTP Secret can't be read from DB.
	if errOTP != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errOTP.Error()})
		return
	}

	// Compare submitted code to OTP Secret.
	code := c.GetString("code")
	valid := security.CheckCode(code, OTP.Totpsecret)

	// Abort request when OTP code seems invalid.
	if valid == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid code"})
		return
	}

	c.Next()
}

func (e *Env) MWwriteSessionIntoRedis(c *gin.Context) {
	email := c.GetString("email")

	person, personErr := data.FindPersonByEmail(c.Request.Context(), e.db, email)
	if personErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": personErr.Error()})
		return
	}

	// Deny access when person is inactive.
	if person.Active == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Person inactive."})
		return
	}

	company, companyErr := data.FindCompanyById(c.Request.Context(), e.db, person.CompanyID)
	if companyErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": companyErr.Error()})
		return
	}

	ticket := new(security.Ticket).Create(person, company)

	// Save Ticket.Info in RedisJSON
	status, rjErr := e.cacheJSON.JSONSet(ticket.Stub.String(), ".", ticket.Info)

	if rjErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": rjErr.Error()})
		return
	}

	// Set expiration on Ticket in Redis. The RedisJSON client lacks an API for expirations, but the vanilla Redis client can handle it.
	expire := e.cache.Expire(c.Request.Context(), ticket.Stub.String(), 4*time.Hour)
	if expire.Val() == false {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to set expiration on stub."})
		return
	}

	// After Ticket.Info is written to Redis, then create a session.
	if status.(string) == "OK" {
		session, sessErr := e.cookieShop.Get(c.Request, "sesh")
		if sessErr != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": sessErr.Error()})
			return
		}

		// Assign Ticket.Stub to the session.
		session.Values["stub"] = ticket.Stub.String()

		// Save the session into the *gin.Context, so it can be returned to the web-client.
		saveErr := session.Save(c.Request, c.Writer)
		if saveErr != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": saveErr.Error()})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to set SessionInfo."})
		return
	}

	c.Status(http.StatusOK)
}
