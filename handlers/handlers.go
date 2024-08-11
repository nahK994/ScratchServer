package handlers

import (
	"fmt"
	"net"
	"strings"

	"github.com/nahK994/ScratchServer/models"
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

func HandleResponse(conn net.Conn) {
	statusCode := 200
	statusText := "OK"
	responseBody := "Hello, World!"
	contentLength := len(responseBody)
	contentType := "text/plain"
	// contentType := "application/json"

	response := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\n"+
			"Content-Type: %s\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s",
		statusCode, statusText, contentType, contentLength, responseBody,
	)
	fmt.Println(response)
	conn.Write([]byte(response))
}
