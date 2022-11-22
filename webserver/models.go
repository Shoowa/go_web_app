package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/models"
	"github.com/gin-gonic/gin"
	null "github.com/volatiletech/null/v8"
)

type requestPostModel struct {
	Code       string    `json:"code" binding:"required,alphanumunicode"`
	CommonID   int       `json:"common_id" binding:"required,numeric"`
	CategoryID int       `json:"category_id" binding:"required,numeric"`
	BrandID    int       `json:"brand_id" binding:"required,numeric"`
	Traits     null.JSON `json:"traits" binding:"required,json"`
}

func (e *Env) CreateModelPOST(c *gin.Context) {
	var reqModel requestPostModel
	if reqErr := c.ShouldBindJSON(&reqModel); reqErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Transform request into ORM-model.
	model := models.Model{
		Code:       reqModel.Code,
		CommonID:   reqModel.CommonID,
		CategoryID: reqModel.CategoryID,
		BrandID:    reqModel.BrandID,
		Traits:     reqModel.Traits,
	}

	dbErr := data.CreateModel(c.Request.Context(), e.db, model)

	if dbErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	c.Status(http.StatusOK)
}

type requestGetModels struct {
	CategoryID int `uri:"categoryID" binding:"required,numeric"`
	BrandID    int `uri:"brandID" binding:"required,numeric"`
}

func (e *Env) FindModelsByCategoryIDAndBrandIDGET(c *gin.Context) {
	var reqModel requestGetModels
	if reqErr := c.ShouldBindUri(&reqModel); reqErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	models, err := data.FindModelsByCategoryIDAndBrandID(c.Request.Context(), e.db, reqModel.CategoryID, reqModel.BrandID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	if models == nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"msg": "No items available."})
		return
	}

	c.IndentedJSON(http.StatusOK, models)
}
