package webserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	data "github.com/Shoowa/broker.git/database"
)

func (suite *Routing) TestCreateCompanyPOST() {
	payload := bytes.NewBuffer(data.RandomCompanyJSON())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v0/signup", payload)
	suite.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}
