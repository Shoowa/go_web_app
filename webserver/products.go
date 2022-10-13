package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/models"
	"github.com/gin-gonic/gin"
)

func (e *Env) ProductsGET(c *gin.Context) {
	catalog, err := data.ReadProducts(c.Request.Context(), e.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "bad"})
	}

	c.JSON(http.StatusOK, catalog)
}

func (e *Env) ProductsFromCompanyGET(c *gin.Context) {
	name := c.Param("company")

	catalog, err := data.ReadProductsFromCompany(c.Request.Context(), e.db, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "bad"})
		return
	}

	if catalog == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.IndentedJSON(http.StatusOK, catalog)
}

type Identity struct {
	ID int `uri:"id", binding:"required"`
}

func (e *Env) ProductsByCompanyIdGET(c *gin.Context) {
	var input Identity
	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	catalog, err := data.FindProductsByCompanyId(c.Request.Context(), e.db, input.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "bad"})
		return
	}

	if catalog == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.IndentedJSON(http.StatusOK, catalog)
}

func (e *Env) CreateProductPOST(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.CreateProduct(c.Request.Context(), e.db, product)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
