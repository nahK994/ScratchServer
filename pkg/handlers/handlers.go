package handlers

import (
	"fmt"
	"net/http"

	"github.com/nahK994/SimpleServer/pkg/errors"
	"github.com/nahK994/SimpleServer/pkg/models"
	"github.com/nahK994/SimpleServer/pkg/utils"
)

func getRequestHandler(urlPath models.HttpUrlPath, req *models.Request) (models.HttpHandlerFunc, error) {
	var err error = nil
	var handleFunc models.HttpHandlerFunc = nil

	requestHandlers, ok := utils.HttpRouteMapper[urlPath]
	if !ok {
		err = errors.UrlNotFound{}
	}

	for _, item := range requestHandlers {
		if req.Method == item.Method {
			handleFunc = item.Func
			break
		}
	}

	if handleFunc == nil {
		err = errors.MethodNotAllowed{}
	}

	return handleFunc, err
}

func handleError(err error, res *models.Response) {
	switch err.(type) {
	case errors.UrlNotFound:
		res.StatusCode = http.StatusNotFound
	case errors.MethodNotAllowed:
		res.StatusCode = http.StatusMethodNotAllowed
	}
	res.Body = err.Error()
}

func HandleRequest(msg []byte) *models.Response {
	req := ParseHttpRequest(msg)
	res := new(models.Response)

	requestHandler, err := getRequestHandler(models.HttpUrlPath(req.UrlPath), req)
	if err != nil {
		handleError(err, res)
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
