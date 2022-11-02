package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const dot = "."

func (e *Env) AuthN(c *gin.Context) {
	session, sessErr := e.cookieShop.Get(c.Request, "sesh")

	if sessErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": sessErr.Error()})
		return
	}

	// Read a ticket stub from the client's session.
	value, ok := session.Values["stub"]
	if ok == false {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied, because stub is missing."})
		return
	}
	stub := value.(string)

	// Do we have a record of that stub in the Redis cache?
	exist, existErr := e.cache.Exists(c.Request.Context(), stub).Result()

	if existErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": existErr.Error()})
		return
	}

	// When we lack a record of that stub, deny access to the client.
	if exist == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error:": "Denied!"})
		return
	}

	// When we have a record of that stub, admit the client.
	c.Set("stub", stub)
	c.Next()
}
