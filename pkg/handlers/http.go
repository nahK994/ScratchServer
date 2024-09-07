package handlers

import (
	"strings"

	"github.com/nahK994/SimpleServer/pkg/models"
)

func ParseHttpRequest(req []byte) *models.Request {
	cmdLines := strings.Split(string(req), "\r\n")
	aa := strings.Split(cmdLines[0], " ")
	return &models.Request{
		Method:  aa[0],
		UrlPath: aa[1],
		Body:    cmdLines[len(cmdLines)-1],
	}
}
