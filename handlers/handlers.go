package handlers

import (
	"fmt"
	"net/http"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

func HandleRequest(msg []byte) *models.Response {
	req := ParseHttpRequest(msg)
	response := new(models.Response)
	requestHandler, foundUrlPath := utils.HttpRouteMapper[models.HttpUrlPath(req.UrlPath)]
	if !foundUrlPath {
		response.StatusCode = http.StatusNotFound
		response.Body = utils.StatusText[404]
	} else if requestHandler.Method != req.Method {
		response.StatusCode = 405
		response.Body = utils.StatusText[405]
	} else {
		requestHandler.Func(*req, response)
	}
	return response
}

func HandleResponse(response *models.Response) string {
	statusCode := response.StatusCode
	statusText := utils.StatusText[statusCode]
	contentType, responseBody := extractResponseBody(response.Body)
	contentLength := len(responseBody)

	resp := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\n"+
			"Content-Type: %s\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s",
		statusCode, statusText, contentType, contentLength, responseBody,
	)
	// fmt.Println(resp)
	return resp
}
