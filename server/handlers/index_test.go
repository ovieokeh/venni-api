package venni

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "github.com/ovieokeh/venni-api/server/utils"
	"gotest.tools/assert"
)

func TestIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	handler.ServeHTTP(rr, req)

	expectedStatusCode := 200
	expectedMessage := "welcome to venni 1.0"
	var parsedResponse utils.HTTPResponse

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &parsedResponse)

	assert.Equal(t, expectedStatusCode, rr.Code)
	assert.Equal(t, expectedMessage, parsedResponse.Message)
}

func TestNotfound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/lost", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NotFoundHandler)

	handler.ServeHTTP(rr, req)

	expectedStatusCode := 404
	expectedMessage := "endpoint not found"
	var parsedResponse utils.HTTPResponse

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &parsedResponse)

	assert.Equal(t, expectedStatusCode, rr.Code)
	assert.Equal(t, expectedMessage, parsedResponse.Message)
}
