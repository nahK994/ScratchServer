package handlers

import (
	"fmt"
	"net/http"

	"github.com/nahK994/TCPickle/pkg/errors"
	"github.com/nahK994/TCPickle/pkg/models"
	"github.com/nahK994/TCPickle/pkg/utils"
)

func getRequestHandler(urlPath models.HttpUrlPath, req *models.Request) (models.HttpHandlerFunc, error) {
	requestHandler, ok := utils.HttpRouteMapper[urlPath]
	var err error = nil
	if !ok {
		err = errors.UrlNotFound{}
	} else if requestHandler.Method != req.Method {
		err = errors.MethodNotAllowed{}
	}
	return requestHandler.Func, err
}

func HandleRequest(msg []byte) *models.Response {
	req := ParseHttpRequest(msg)
	res := new(models.Response)

	requestHandler, err := getRequestHandler(models.HttpUrlPath(req.UrlPath), req)
	if err != nil {
		switch err.(type) {
		case errors.UrlNotFound:
			res.StatusCode = http.StatusNotFound
			res.Body = err.Error()
		case errors.MethodNotAllowed:
			res.StatusCode = http.StatusMethodNotAllowed
			res.Body = err.Error()
		}
	} else {
		requestHandler(*req, res)
	}

	return res
}

func HandleResponse(response *models.Response) string {
	statusCode := response.StatusCode
	statusText := utils.StatusText[statusCode]
	contentType := "application/json"
	responseBody, _ := response.Body.(string)
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
