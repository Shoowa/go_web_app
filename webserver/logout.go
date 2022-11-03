package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e *Env) SessionDEL(c *gin.Context) {
	stub := c.GetString("stub")

	// Delete ticket info from the Redis cache.
	ok, err := e.cache.Del(c.Request.Context(), stub).Result()

	// Explain error.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Explain failure to delete when key is missing.
	if ok == 0 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Session information destroyed prior to request."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Goodbye!"})
}
