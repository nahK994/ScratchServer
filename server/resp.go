package server

import (
	"github.com/nahK994/TCPickle/models"
)

type RespServer struct {
	Config
}

func (s *RespServer) RequestHandler(handler models.HttpHandlerFunc) {

}
