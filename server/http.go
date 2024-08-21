package server

import (
	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

type HttpServer struct {
	Config
}

func (s *HttpServer) RequestHandler(urlPath models.HttpUrlPath, method string, handler models.HttpHandlerFunc) {
	utils.HttpRouteMapper[urlPath] = models.HttpHandler{
		Method: method,
		Func:   handler,
	}
}
