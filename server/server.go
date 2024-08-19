package server

import (
	"log/slog"
	"net"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

type Server struct {
	ListenAddress string
	ln            net.Listener
}

func Initiate(listenAddress string) *Server {
	return &Server{
		ListenAddress: listenAddress,
	}
}

func (s *Server) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := NewPeer(conn)
		go peer.readConn()
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		return err
	}

	s.ln = ln
	defer ln.Close()

	slog.Info("server running", "listenAddr", s.ListenAddress)

	return s.acceptConn()
}

func (s *Server) RequestHandler(urlPath models.HttpUrlPath, method string, handler models.HandleHttpFunc) {
	utils.RouteMapper[urlPath] = models.HandlerDetails{
		Method: method,
		Func:   handler,
	}
}
