package handlers

import (
	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

func HandleRequest(req []byte, protocol string) (*models.HttpResponse, *models.RespResponse) {
	if protocol == utils.HTTP {
		request := ParseHttpRequest(req)
		response := new(models.HttpResponse)
		requestHandler := utils.RouteMapper[models.HttpUrlPath(request.UrlPath)]
		if requestHandler.Method != request.Method {
			response.StatusCode = 405
			response.Body = ""
		} else {
			requestHandler.Func(*request, response)
		}
		return response, nil
	} else if protocol == utils.RESP {

	}

	return nil, nil
}

func HandleResponse(httpRes *models.HttpResponse, respRes *models.RespResponse) string {
	if httpRes != nil {
		return HandleHttpResponse(httpRes)
	} else if respRes != nil {
		return HandleRespResponse(respRes)
	}
	return ""
}
