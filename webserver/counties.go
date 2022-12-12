package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
)

func (e *Env) FindCountiesByStateIDGET(c *gin.Context) {
	var input State
	if errInput := c.ShouldBindUri(&input); errInput != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errInput.Error()})
		return
	}

	counties, err := data.FindCountiesByStateID(c.Request.Context(), e.db, input.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, counties)
}
