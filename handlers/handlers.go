package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

type content_type string
type response_body string

func extractResponseBody(body interface{}) (content_type, response_body) {
	var contentType content_type
	var responseBody response_body

	if body1, ok1 := body.(string); ok1 {
		contentType = "text/plain"
		responseBody = response_body(body1)
	} else if body2, ok2 := json.Marshal(body); ok2 == nil {
		contentType = "application/json"
		responseBody = response_body(body2)
	} else {
		// TODO: implement FUCK
	}

	return contentType, responseBody
}

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
	conn.Write([]byte(resp))
}
