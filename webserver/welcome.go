package webserver

import (
	"fmt"
	"net/http"

	"github.com/Shoowa/broker.git/cache"
	"github.com/gin-gonic/gin"
)

func (e *Env) WelcomeGET(c *gin.Context) {
	stub := c.GetString("stub")

	email, err := cache.ReadSessionEmail(e.cacheJSON, stub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Hi, %s!", *email)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}
