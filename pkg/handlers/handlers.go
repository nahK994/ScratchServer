package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nahK994/SimpleServer/pkg/errors"
	"github.com/nahK994/SimpleServer/pkg/types"
	"github.com/nahK994/SimpleServer/pkg/utils"
)

func parseHttpRequest(req []byte) *types.Request {
	cmdLines := strings.Split(string(req), "\r\n")
	aa := strings.Split(cmdLines[0], " ")
	return &types.Request{
		Method:  aa[0],
		UrlPath: aa[1],
		Body:    cmdLines[len(cmdLines)-1],
	}
}

func getRequestHandler(urlPath types.HttpUrlPath, req *types.Request) (types.HttpHandlerFunc, error) {
	var err error = nil
	var handleFunc types.HttpHandlerFunc = nil

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

func handleError(err error, res *types.Response) {
	switch err.(type) {
	case errors.UrlNotFound:
		res.StatusCode = http.StatusNotFound
	case errors.MethodNotAllowed:
		res.StatusCode = http.StatusMethodNotAllowed
	}
	res.Body = err.Error()
}

func HandleRequest(msg []byte) *types.Response {
	req := parseHttpRequest(msg)
	res := new(types.Response)

	requestHandler, err := getRequestHandler(types.HttpUrlPath(req.UrlPath), req)
	if err != nil {
		handleError(err, res)
	} else {
		requestHandler(*req, res)
	}

	return res
}

func HandleResponse(response *types.Response) string {
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
