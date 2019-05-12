package venni

import (
	"net/http"

	utils "github.com/ovieokeh/venni-api/server/utils"
)

// IndexHandler - handles calls to "/"
func IndexHandler(Response http.ResponseWriter, Request *http.Request) {
	response := utils.HTTPResponse{
		Status:  "success",
		Message: "welcome to venni 1.0",
		Data:    nil,
	}

	utils.RespondWithJSON(Response, 200, response)
}

// NotFoundHandler - handles calls to any unmatched route
func NotFoundHandler(Response http.ResponseWriter, Request *http.Request) {
	response := utils.HTTPResponse{
		Status:  "error",
		Message: "endpoint not found",
		Data:    nil,
	}

	utils.RespondWithJSON(Response, 404, response)
}
