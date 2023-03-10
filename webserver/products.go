package webserver

import (
	"net/http"

	"github.com/Shoowa/broker.git/cache"
	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/models"
	"github.com/gin-gonic/gin"
)

func (e *Env) NotMyProductsGET(c *gin.Context) {
	stub := c.GetString("stub")

	myCompany, cacheErr := cache.ReadSessionCompany(e.cacheJSON, stub)
	if cacheErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": cacheErr.Error()})
		return
	}

	catalog, dbErr := data.FindProductsExcludeCompanyId(c.Request.Context(), e.db, *myCompany)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": dbErr.Error()})
		return
	}

	if catalog == nil {
		c.JSON(http.StatusNoContent, gin.H{"msg": "No items listed."})
		return
	}

	c.IndentedJSON(http.StatusOK, catalog)
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

func (e *Env) MyProductsGET(c *gin.Context) {
	stub := c.GetString("stub")

	myCompany, cacheErr := cache.ReadSessionCompany(e.cacheJSON, stub)
	if cacheErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": cacheErr.Error()})
		return
	}

	catalog, dbErr := data.FindProductsByCompanyId(c.Request.Context(), e.db, *myCompany)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	if catalog == nil {
		c.JSON(http.StatusNoContent, gin.H{"msg": "No items listed."})
		return
	}

	c.IndentedJSON(http.StatusOK, catalog)
}

func (e *Env) FindActiveProductsByModelCodeGET(c *gin.Context) {
	modelCode := c.Param("code")

	catalog, err := data.FindActiveProductsByModelCode(c.Request.Context(), e.db, modelCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	if catalog == nil {
		c.JSON(http.StatusNoContent, gin.H{"msg": "No items available."})
		return
	}

	c.IndentedJSON(http.StatusOK, catalog)
}
