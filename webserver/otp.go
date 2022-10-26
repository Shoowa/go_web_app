package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/security"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp"
)

type requestedEmail struct {
	Email string `json:"email" binding:"required,email"`
}

func (e *Env) MWreadEmail(c *gin.Context) {
	var identity requestedEmail
	if err := c.ShouldBindJSON(&identity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Set("email", identity.Email)
	c.Next()
}

func (e *Env) MWisEmailVerified(c *gin.Context) {
	personEmail := c.GetString("email")

	// Find ID by email
	person, err := data.FindPersonByEmail(c.Request.Context(), e.db, personEmail)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Is this email verified?
	verified, err2 := data.IsEmailVerified(c.Request.Context(), e.db, person.ID)
	if err2 != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Abort request when e-mail remains unverified.
	if verified == false {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Set CONTEXT true
	c.Next()
}

func (e *Env) MWdoYouHaveOTP(c *gin.Context) {
	personEmail := c.GetString("email")
	exist, err := data.OtpSecretExists(c.Request.Context(), e.db, personEmail)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if exist == true {
		c.AbortWithStatusJSON(http.StatusSeeOther, gin.H{"status": "OTP Secret already exists."})
		return
	}

	c.Next()
}

func (e *Env) MWcreateOTP(c *gin.Context) {
	personEmail := c.GetString("email")

	key, keyErr := security.GenerateKey(personEmail)
	if keyErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	creationErr := data.CreateOtpKey(c.Request.Context(), e.db, personEmail, key.Secret())
	if creationErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("otpKey", key)
	c.Next()
}

func (e *Env) MWshowOTP(c *gin.Context) {
	keyString, exists := c.Get("otpKey")
	if exists == false {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	key := keyString.(*otp.Key)
	img, qrErr := security.CreateQR(key)
	if qrErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	buffer := security.BufferQR(img)

	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())
	c.Next()
}
