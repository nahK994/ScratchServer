package server

import (
	"log/slog"
	"net"

	"github.com/nahK994/TCPickle/models"
	"github.com/nahK994/TCPickle/utils"
)

type Config struct {
	listenAddress string
	ln            net.Listener
}

func Initiate(listenAddress string) *Config {
	return &Config{
		listenAddress: listenAddress,
	}
}

func (s *Config) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := NewPeer(conn)
		go peer.readConn()
	}
}

func (s *Config) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}

	s.ln = ln
	defer ln.Close()

	slog.Info("server running", "listenAddr", s.listenAddress)

	return s.acceptConn()
}

func (s *Config) RequestHandler(urlPath models.HttpUrlPath, method string, handler models.HttpHandlerFunc) {
	utils.HttpRouteMapper[urlPath] = models.HttpHandler{
		Method: method,
		Func:   handler,
	}
}
