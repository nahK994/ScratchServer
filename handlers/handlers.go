package handlers

import (
	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

func HandleRequest(req []byte, protocol string) *models.Response {
	if protocol == utils.HTTP {
		request := ParseHttpRequest(req)
		response := new(models.HttpResponse)
		requestHandler, foundUrlPath := utils.RouteMapper[models.HttpUrlPath(request.UrlPath)]
		if !foundUrlPath {
			response.StatusCode = 404
			response.Body = utils.StatusText[404]
		} else if requestHandler.Method != request.Method {
			response.StatusCode = 405
			response.Body = utils.StatusText[405]
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
