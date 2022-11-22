package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
)

func (e *Env) FindAllBrandsGET(c *gin.Context) {
	list, err := data.FindAllBrands(c.Request.Context(), e.db)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, list)
}
