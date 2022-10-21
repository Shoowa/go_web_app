package webserver

import (
	"net/http"

	data "github.com/Shoowa/broker.git/database"
	"github.com/Shoowa/broker.git/models"
	"github.com/Shoowa/broker.git/security"
	"github.com/gin-gonic/gin"
	null "github.com/volatiletech/null/v8"
)

type requestCreateCompany struct {
	Name      string      `json:"name" binding:"required"`
	Email     string      `json:"email" binding:"required,email"`
	Telephone string      `json:"telephone" binding:"required,e164=US"`
	Ext       null.String `json:"ext"`
	Street    string      `json:"street" binding:"required"`
	Suite     null.String `json:"suite"`
	State     string      `json:"state" binding:"required"`
	County    int         `json:"county" binding:"required"`
	City      null.Int    `json:"city"`
	Postal    string      `json:"postal", binding:"required,postcode_iso3166_alpha2=US"`
	Org       int         `json:"org" binding:"required"`
}

type requestCreatePerson struct {
	Email    string `json:"email" binding:"required,email"`
	Title    string `json:"title"`
	Forename string `json:"forename" binding:"required,alpha"`
	Surname  string `json:"surname" binding:"required,alpha"`
}

type requestCreatePW struct {
	PW string `json:"pw" binding:"required,gte=12,containsany=!@#$"`
}

type Combination struct {
	Company  requestCreateCompany
	Person   requestCreatePerson
	Password requestCreatePW
}

func (e *Env) SignupCompanyPOST(c *gin.Context) {
	c.Status(http.StatusOK)
}

// Read JSON Payload's sub-records.
func (e *Env) MWreadSubmission(c *gin.Context) {
	var combo Combination
	if err := c.ShouldBindJSON(&combo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("company", combo.Company)
	c.Set("person", combo.Person)
	c.Set("password", combo.Password)
	c.Next()
}

func (e *Env) MWregisterCompany(c *gin.Context) {
	company, ok := c.Get("company")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing company info for creation"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	companyInfo := company.(requestCreateCompany)

	freshCompany := models.Company{
		Name:      companyInfo.Name,
		Email:     companyInfo.Email,
		Telephone: companyInfo.Telephone,
		Ext:       companyInfo.Ext,
		Street:    companyInfo.Street,
		Suite:     companyInfo.Suite,
		Postal:    companyInfo.Postal,
		StateID:   companyInfo.State,
		CountyID:  companyInfo.County,
		MuniID:    companyInfo.City,
		StructID:  companyInfo.Org,
	}

	err := data.CreateCompany(c.Request.Context(), e.db, freshCompany)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Next()
}

func (e *Env) MWfindCompanyID(c *gin.Context) {
	company, ok := c.Get("company")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing company info for selection"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	companyInfo := company.(requestCreateCompany)

	// query company by name for ID
	companyID, err := data.FindCompanyIDByName(c.Request.Context(), e.db, companyInfo.Name, companyInfo.State)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("companyID", companyID)
	c.Next()
}

func (e *Env) MWregisterPerson(c *gin.Context) {
	person, ok := c.Get("person")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing person info"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	personInfo := person.(requestCreatePerson)

	companyID, ok := c.Get("companyID")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing companyID info"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	freshPerson := models.Person{
		CompanyID: companyID.(int),
		Email:     personInfo.Email,
		Title:     personInfo.Title,
		Forename:  personInfo.Forename,
		Surname:   personInfo.Surname,
	}

	personErr := data.CreatePerson(c.Request.Context(), e.db, freshPerson)
	if personErr != nil {
		c.Status(http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("email", personInfo.Email)
	c.Next()
}

func (e *Env) MWdigestPW(c *gin.Context) {
	pw, ok := c.Get("password")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing password info"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	reqPW := pw.(requestCreatePW)

	hash, err := security.HashPW(reqPW.PW)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("pwdigest", hash)
	c.Next()
}

func (e *Env) MWregisterPWDigest(c *gin.Context) {
	pwDigest, ok := c.Get("pwdigest")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing pwdigest info"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	email, ok := c.Get("email")
	if ok == false {
		c.JSON(http.StatusNoContent, gin.H{"msg": "missing email info"})
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	freshDigest := models.Digest{
		Email:  email.(string),
		Digest: pwDigest.(string),
	}

	digestErr := data.CreateDigest(c.Request.Context(), e.db, freshDigest)
	if digestErr != nil {
		c.Status(http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Next()
}
