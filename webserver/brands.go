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

func (e *Env) FindBrandsRelatedToModelIDGET(c *gin.Context) {
	var input Identity
	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eligibleBrands, err := data.FindBrandsRelatedToModelID(c.Request.Context(), e.db, input.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	if eligibleBrands == nil {
		c.JSON(http.StatusNoContent, gin.H{"msg": "No items available."})
		return
	}

	c.IndentedJSON(http.StatusOK, eligibleBrands)
}
