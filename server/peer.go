package server

import (
	"log/slog"
	"net"

	"github.com/nahK994/TCPickle/handlers"
)

type Peer struct {
	conn net.Conn
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
	}
}

func (p *Peer) readConn() {
	buf := make([]byte, 1024)
	n, err := p.conn.Read(buf)
	if err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", p.conn.RemoteAddr())
		p.conn.Close()
		return
	}

	response := handlers.HandleRequest(buf[:n])
	resp := handlers.HandleResponse(response)
	p.conn.Write([]byte(resp))
	p.conn.Close()
}
