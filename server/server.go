package server

import (
	"log/slog"
	"net"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

type Server struct {
	listenAddress string
	ln            net.Listener
	protocol      string
}

func Initiate(listenAddress, protocol string) *Server {
	return &Server{
		listenAddress: listenAddress,
		protocol:      protocol,
	}
}

func (s *Server) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := NewPeer(conn)
		go peer.readConn(s.protocol)
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	s.ln = ln
	defer ln.Close()

	slog.Info("server running", "listenAddr", s.listenAddress)

	return s.acceptConn()
}

func (s *Server) RequestHandler(urlPath models.HttpUrlPath, method string, handler models.HandleHttpFunc) {
	utils.RouteMapper[urlPath] = models.HandlerDetails{
		Method: method,
		Func:   handler,
	}
}
