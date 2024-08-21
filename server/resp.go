package server

import (
	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

type RespServer struct {
	Config
}

func (s *RespServer) RequestHandler(handler models.RespHandlerFunc) {
	utils.RespHandleFunc = handler
}
