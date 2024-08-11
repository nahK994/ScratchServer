package handlers

import (
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

	// fmt.Println("1...", cmdLines[0])
	// fmt.Println("2...", cmdLines[len(cmdLines)-1])
	// var test LoginReq
	// json.Unmarshal([]byte(cmdLines[len(cmdLines)-1]), &test)
	// fmt.Println("Final ==> ", test)
}
