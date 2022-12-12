package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
)

type State struct {
	ID string `uri:"id", binding:"required"`
}

func (e *Env) FindAllStatesGET(c *gin.Context) {
	states, err := data.FindAllStates(c.Request.Context(), e.db)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, states)
}
