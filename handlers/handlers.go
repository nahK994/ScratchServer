package handlers

import (
	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

func HandleRequest(req []byte, protocol string) *models.Response {
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
		return &models.Response{
			Http: *response,
		}
	} else if protocol == utils.RESP {

	}

	return nil
}

func HandleResponse(res *models.Response, protocol string) string {
	if protocol == utils.HTTP {
		return HandleHttpResponse(&res.Http)
	} else if protocol == utils.RESP {
		return HandleRespResponse(&res.Resp)
	}
	return ""
}
