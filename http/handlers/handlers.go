package handlers

import (
	"strings"

	"github.com/nahK994/ScratchServer/http/utils"
)

func HandleRequest(bufStr string) *utils.Request {
	aa := strings.Split(bufStr, " ")
	method := aa[0]
	urlPath := aa[1]
	return &utils.Request{
		Method:  method,
		UrlPath: urlPath,
	}
}
