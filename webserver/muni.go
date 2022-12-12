package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
)

func (e *Env) FindMunisByCountyIDGET(c *gin.Context) {
	var input Identity
	if errInput := c.ShouldBindUri(&input); errInput != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errInput.Error()})
		return
	}

	munis, err := data.FindMuniByCountyID(c.Request.Context(), e.db, input.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	if munis == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.IndentedJSON(http.StatusOK, munis)
}
