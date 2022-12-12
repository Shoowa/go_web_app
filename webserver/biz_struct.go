package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
)

func (e *Env) FindAllOrgsGET(c *gin.Context) {
	orgs, err := data.FindAllOrgs(c.Request.Context(), e.db)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, orgs)
}
