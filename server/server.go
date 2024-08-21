package server

import (
	"log/slog"
	"net"

	"github.com/nahK994/TCPickle/utils"
)

type Config struct {
	listenAddress string
	ln            net.Listener
	protocol      string
}

func InitiateHttp(listenAddress string) *HttpServer {
	return &HttpServer{
		Config{
			listenAddress: listenAddress,
			protocol:      utils.HTTP,
		},
	}
}

func InitiateResp(listenAddress string) *RespServer {
	return &RespServer{
		Config{
			listenAddress: listenAddress,
			protocol:      utils.RESP,
		},
	}
}

func (s *Config) acceptConn() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			return err
		}

		peer := NewPeer(conn)
		go peer.readConn(s.protocol)
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
