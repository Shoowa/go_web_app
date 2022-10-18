package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shoowa/broker.git/cache"
	data "github.com/Shoowa/broker.git/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type Routing struct {
	suite.Suite
	Router *gin.Engine
}

func (suite *Routing) SetupSuite() {
	c := data.DefineConfig()
	db := c.Access()

	red := cache.Setup()
	redjson := cache.SetupRedisJSONClient(red)

	suite.Router = NewRouter(db, red, redjson)
}

func TestRouting(t *testing.T) {
	suite.Run(t, new(Routing))
}

func (suite *Routing) TestHealthGET() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}
