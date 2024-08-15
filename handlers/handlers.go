package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/nahK994/ScratchServer/models"
	"github.com/nahK994/ScratchServer/utils"
)

func HandleRequest(req []byte) *models.Request {
	cmdLines := strings.Split(string(req), "\r\n")
	aa := strings.Split(cmdLines[0], " ")
	return &models.Request{
		Method:  aa[0],
		UrlPath: aa[1],
		Body:    cmdLines[len(cmdLines)-1],
	}
}

func HandleResponse(response *models.Response, conn net.Conn) {
	// var data interface{} = map[string]interface{}{
	// 	"name": "khan",
	// }
	// response.StatusCode = 201
	// response.Body = data

	statusCode := response.StatusCode
	statusText := utils.StatusText[statusCode]

	var contentLength int
	var contentType string
	var responseBody string
	if body1, ok1 := response.Body.(string); ok1 {
		contentType = "text/plain"
		responseBody = body1
	} else if body2, ok2 := json.Marshal(response.Body); ok2 == nil {
		contentType = "application/json"
		responseBody = string(body2)
	} else {
		// TODO: implement FUCK
	}
	contentLength = len(responseBody)
	resp := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\n"+
			"Content-Type: %s\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s",
		statusCode, statusText, contentType, contentLength, responseBody,
	)
	// fmt.Println(resp)
	conn.Write([]byte(resp))
}
