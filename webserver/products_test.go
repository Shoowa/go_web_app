package webserver

import (
	"net/http"
	"net/http/httptest"
)

func (suite *Routing) TestExistingProductsFromCompanyGET() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v0/products/Cheap Cups", nil)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *Routing) TestAbsentProductsFromCompanyGET() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v0/products/CheapXCups", nil)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusNoContent, w.Code)
}
