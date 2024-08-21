package handlers

import (
	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

func HandleRequest(msg []byte, protocol string) *models.Response {
	if protocol == utils.HTTP {
		req := ParseHttpRequest(msg)
		response := new(models.HttpResponse)
		requestHandler, foundUrlPath := utils.HttpRouteMapper[models.HttpUrlPath(req.UrlPath)]
		if !foundUrlPath {
			response.StatusCode = 404
			response.Body = utils.StatusText[404]
		} else if requestHandler.Method != req.Method {
			response.StatusCode = 405
			response.Body = utils.StatusText[405]
		} else {
			requestHandler.Func(*req, response)
		}
		return &models.Response{
			Http: *response,
		}
	} else if protocol == utils.RESP {
		req := models.RespRequest{
			Request: string(msg),
		}
		res := new(models.RespResponse)
		utils.RespHandleFunc(req, res)
		return &models.Response{
			Resp: *res,
		}
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
